/*
Copyright 2024 The ORC Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PortResourceStatusApplyConfiguration represents a declarative configuration of the PortResourceStatus type for use
// with apply.
type PortResourceStatusApplyConfiguration struct {
	Name                                    *string                                      `json:"name,omitempty"`
	Description                             *string                                      `json:"description,omitempty"`
	NetworkID                               *string                                      `json:"networkID,omitempty"`
	ProjectID                               *string                                      `json:"projectID,omitempty"`
	Status                                  *string                                      `json:"status,omitempty"`
	Tags                                    []string                                     `json:"tags,omitempty"`
	AdminStateUp                            *bool                                        `json:"adminStateUp,omitempty"`
	MACAddress                              *string                                      `json:"macAddress,omitempty"`
	DeviceID                                *string                                      `json:"deviceID,omitempty"`
	DeviceOwner                             *string                                      `json:"deviceOwner,omitempty"`
	AllowedAddressPairs                     []AllowedAddressPairStatusApplyConfiguration `json:"allowedAddressPairs,omitempty"`
	FixedIPs                                []FixedIPStatusApplyConfiguration            `json:"fixedIPs,omitempty"`
	SecurityGroups                          []string                                     `json:"securityGroups,omitempty"`
	PropagateUplinkStatus                   *bool                                        `json:"propagateUplinkStatus,omitempty"`
	VNICType                                *string                                      `json:"vnicType,omitempty"`
	PortSecurityEnabled                     *bool                                        `json:"portSecurityEnabled,omitempty"`
	NeutronStatusMetadataApplyConfiguration `json:",inline"`
}

// PortResourceStatusApplyConfiguration constructs a declarative configuration of the PortResourceStatus type for use with
// apply.
func PortResourceStatus() *PortResourceStatusApplyConfiguration {
	return &PortResourceStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithName(value string) *PortResourceStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithDescription(value string) *PortResourceStatusApplyConfiguration {
	b.Description = &value
	return b
}

// WithNetworkID sets the NetworkID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkID field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithNetworkID(value string) *PortResourceStatusApplyConfiguration {
	b.NetworkID = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithProjectID(value string) *PortResourceStatusApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithStatus(value string) *PortResourceStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *PortResourceStatusApplyConfiguration) WithTags(values ...string) *PortResourceStatusApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithAdminStateUp sets the AdminStateUp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminStateUp field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithAdminStateUp(value bool) *PortResourceStatusApplyConfiguration {
	b.AdminStateUp = &value
	return b
}

// WithMACAddress sets the MACAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MACAddress field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithMACAddress(value string) *PortResourceStatusApplyConfiguration {
	b.MACAddress = &value
	return b
}

// WithDeviceID sets the DeviceID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeviceID field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithDeviceID(value string) *PortResourceStatusApplyConfiguration {
	b.DeviceID = &value
	return b
}

// WithDeviceOwner sets the DeviceOwner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeviceOwner field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithDeviceOwner(value string) *PortResourceStatusApplyConfiguration {
	b.DeviceOwner = &value
	return b
}

// WithAllowedAddressPairs adds the given value to the AllowedAddressPairs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedAddressPairs field.
func (b *PortResourceStatusApplyConfiguration) WithAllowedAddressPairs(values ...*AllowedAddressPairStatusApplyConfiguration) *PortResourceStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllowedAddressPairs")
		}
		b.AllowedAddressPairs = append(b.AllowedAddressPairs, *values[i])
	}
	return b
}

// WithFixedIPs adds the given value to the FixedIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FixedIPs field.
func (b *PortResourceStatusApplyConfiguration) WithFixedIPs(values ...*FixedIPStatusApplyConfiguration) *PortResourceStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFixedIPs")
		}
		b.FixedIPs = append(b.FixedIPs, *values[i])
	}
	return b
}

// WithSecurityGroups adds the given value to the SecurityGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SecurityGroups field.
func (b *PortResourceStatusApplyConfiguration) WithSecurityGroups(values ...string) *PortResourceStatusApplyConfiguration {
	for i := range values {
		b.SecurityGroups = append(b.SecurityGroups, values[i])
	}
	return b
}

// WithPropagateUplinkStatus sets the PropagateUplinkStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PropagateUplinkStatus field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithPropagateUplinkStatus(value bool) *PortResourceStatusApplyConfiguration {
	b.PropagateUplinkStatus = &value
	return b
}

// WithVNICType sets the VNICType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VNICType field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithVNICType(value string) *PortResourceStatusApplyConfiguration {
	b.VNICType = &value
	return b
}

// WithPortSecurityEnabled sets the PortSecurityEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortSecurityEnabled field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithPortSecurityEnabled(value bool) *PortResourceStatusApplyConfiguration {
	b.PortSecurityEnabled = &value
	return b
}

// WithCreatedAt sets the CreatedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreatedAt field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithCreatedAt(value v1.Time) *PortResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.CreatedAt = &value
	return b
}

// WithUpdatedAt sets the UpdatedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedAt field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithUpdatedAt(value v1.Time) *PortResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.UpdatedAt = &value
	return b
}

// WithRevisionNumber sets the RevisionNumber field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionNumber field is set to the value of the last call.
func (b *PortResourceStatusApplyConfiguration) WithRevisionNumber(value int64) *PortResourceStatusApplyConfiguration {
	b.NeutronStatusMetadataApplyConfiguration.RevisionNumber = &value
	return b
}
