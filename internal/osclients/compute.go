/*
Copyright 2021 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package osclients

import (
	"context"
	"fmt"
	"iter"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/attachinterfaces"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/availabilityzones"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servergroups"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/utils/v2/openstack/clientconfig"
)

/*
NovaMinimumMicroversion is the minimum Nova microversion supported by CAPO
2.60 corresponds to OpenStack Queens

For the canonical description of Nova microversions, see
https://docs.openstack.org/nova/latest/reference/api-microversion-history.html

CAPO uses server tags, which were added in microversion 2.52.
CAPO supports multiattach volume types, which were added in microversion 2.60.
*/
const NovaMinimumMicroversion = "2.60"

type ComputeClient interface {
	CreateFlavor(ctx context.Context, opts flavors.CreateOptsBuilder) (*flavors.Flavor, error)
	GetFlavor(ctx context.Context, id string) (*flavors.Flavor, error)
	DeleteFlavor(ctx context.Context, id string) error
	ListFlavors(ctx context.Context, listOpts flavors.ListOptsBuilder) iter.Seq2[*flavors.Flavor, error]

	CreateServer(ctx context.Context, createOpts servers.CreateOptsBuilder, schedulerHints servers.SchedulerHintOptsBuilder) (*servers.Server, error)
	DeleteServer(ctx context.Context, serverID string) error
	GetServer(ctx context.Context, serverID string) (*servers.Server, error)
	ListServers(ctx context.Context, listOpts servers.ListOptsBuilder) iter.Seq2[*servers.Server, error]
}

type computeClient struct{ client *gophercloud.ServiceClient }

// NewComputeClient returns a new compute client.
func NewComputeClient(providerClient *gophercloud.ProviderClient, providerClientOpts *clientconfig.ClientOpts) (ComputeClient, error) {
	compute, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{
		Region:       providerClientOpts.RegionName,
		Availability: clientconfig.GetEndpointType(providerClientOpts.EndpointType),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create compute service client: %v", err)
	}
	compute.Microversion = NovaMinimumMicroversion

	return &computeClient{compute}, nil
}

func (c computeClient) ListAvailabilityZones(ctx context.Context) ([]availabilityzones.AvailabilityZone, error) {
	allPages, err := availabilityzones.List(c.client).AllPages(ctx)
	if err != nil {
		return nil, err
	}
	return availabilityzones.ExtractAvailabilityZones(allPages)
}

func (c computeClient) GetFlavor(ctx context.Context, id string) (*flavors.Flavor, error) {
	return flavors.Get(ctx, c.client, id).Extract()
}

func (c computeClient) CreateFlavor(ctx context.Context, opts flavors.CreateOptsBuilder) (*flavors.Flavor, error) {
	return flavors.Create(ctx, c.client, opts).Extract()
}

func (c computeClient) DeleteFlavor(ctx context.Context, id string) error {
	return flavors.Delete(ctx, c.client, id).ExtractErr()
}

func (c computeClient) ListFlavors(ctx context.Context, opts flavors.ListOptsBuilder) iter.Seq2[*flavors.Flavor, error] {
	pager := flavors.ListDetail(c.client, opts)
	return func(yield func(*flavors.Flavor, error) bool) {
		_ = pager.EachPage(ctx, yieldPage(flavors.ExtractFlavors, yield))
	}
}

func (c computeClient) CreateServer(ctx context.Context, createOpts servers.CreateOptsBuilder, schedulerHints servers.SchedulerHintOptsBuilder) (*servers.Server, error) {
	return servers.Create(ctx, c.client, createOpts, schedulerHints).Extract()
}

func (c computeClient) DeleteServer(ctx context.Context, serverID string) error {
	return servers.Delete(ctx, c.client, serverID).ExtractErr()
}

func (c computeClient) GetServer(ctx context.Context, serverID string) (*servers.Server, error) {
	return servers.Get(ctx, c.client, serverID).Extract()
}

func (c computeClient) ListServers(ctx context.Context, opts servers.ListOptsBuilder) iter.Seq2[*servers.Server, error] {
	pager := servers.List(c.client, opts)
	return func(yield func(*servers.Server, error) bool) {
		_ = pager.EachPage(ctx, yieldPage(servers.ExtractServers, yield))
	}
}

func (c computeClient) ListAttachedInterfaces(ctx context.Context, serverID string) ([]attachinterfaces.Interface, error) {
	interfaces, err := attachinterfaces.List(c.client, serverID).AllPages(ctx)
	if err != nil {
		return nil, err
	}
	return attachinterfaces.ExtractInterfaces(interfaces)
}

func (c computeClient) DeleteAttachedInterface(ctx context.Context, serverID, portID string) error {
	return attachinterfaces.Delete(ctx, c.client, serverID, portID).ExtractErr()
}

func (c computeClient) ListServerGroups(ctx context.Context) ([]servergroups.ServerGroup, error) {
	opts := servergroups.ListOpts{}
	allPages, err := servergroups.List(c.client, opts).AllPages(ctx)
	if err != nil {
		return nil, err
	}
	return servergroups.ExtractServerGroups(allPages)
}

type computeErrorClient struct{ error }

// NewComputeErrorClient returns a ComputeClient in which every method returns the given error.
func NewComputeErrorClient(e error) ComputeClient {
	return computeErrorClient{e}
}
func (e computeErrorClient) CreateFlavor(ctx context.Context, opts flavors.CreateOptsBuilder) (*flavors.Flavor, error) {
	return nil, e.error
}
func (e computeErrorClient) GetFlavor(ctx context.Context, id string) (*flavors.Flavor, error) {
	return nil, e.error
}
func (e computeErrorClient) DeleteFlavor(ctx context.Context, id string) error {
	return e.error
}
func (e computeErrorClient) ListFlavors(_ context.Context, _ flavors.ListOptsBuilder) iter.Seq2[*flavors.Flavor, error] {
	return func(yield func(*flavors.Flavor, error) bool) {
		yield(nil, e.error)
	}
}

func (e computeErrorClient) ListAvailabilityZones(_ context.Context) ([]availabilityzones.AvailabilityZone, error) {
	return nil, e.error
}

func (e computeErrorClient) CreateServer(_ context.Context, _ servers.CreateOptsBuilder, _ servers.SchedulerHintOptsBuilder) (*servers.Server, error) {
	return nil, e.error
}

func (e computeErrorClient) DeleteServer(_ context.Context, _ string) error {
	return e.error
}

func (e computeErrorClient) GetServer(_ context.Context, _ string) (*servers.Server, error) {
	return nil, e.error
}

func (e computeErrorClient) ListServers(ctx context.Context, listOpts servers.ListOptsBuilder) iter.Seq2[*servers.Server, error] {
	return func(yield func(*servers.Server, error) bool) {
		yield(nil, e.error)
	}
}

func (e computeErrorClient) ListAttachedInterfaces(_ context.Context, _ string) ([]attachinterfaces.Interface, error) {
	return nil, e.error
}

func (e computeErrorClient) DeleteAttachedInterface(_ context.Context, _, _ string) error {
	return e.error
}

func (e computeErrorClient) ListServerGroups(_ context.Context) ([]servergroups.ServerGroup, error) {
	return nil, e.error
}
