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
// Code generated by MockGen. DO NOT EDIT.
// Source: ../compute.go
//
// Generated by this command:
//
//	mockgen -package mock -destination=compute.go -source=../compute.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock ComputeClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	iter "iter"
	reflect "reflect"

	flavors "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	servers "github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	gomock "go.uber.org/mock/gomock"
)

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
	isgomock struct{}
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// CreateFlavor mocks base method.
func (m *MockComputeClient) CreateFlavor(ctx context.Context, opts flavors.CreateOptsBuilder) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlavor", ctx, opts)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlavor indicates an expected call of CreateFlavor.
func (mr *MockComputeClientMockRecorder) CreateFlavor(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlavor", reflect.TypeOf((*MockComputeClient)(nil).CreateFlavor), ctx, opts)
}

// CreateServer mocks base method.
func (m *MockComputeClient) CreateServer(ctx context.Context, createOpts servers.CreateOptsBuilder, schedulerHints servers.SchedulerHintOptsBuilder) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", ctx, createOpts, schedulerHints)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockComputeClientMockRecorder) CreateServer(ctx, createOpts, schedulerHints any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockComputeClient)(nil).CreateServer), ctx, createOpts, schedulerHints)
}

// DeleteFlavor mocks base method.
func (m *MockComputeClient) DeleteFlavor(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFlavor", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFlavor indicates an expected call of DeleteFlavor.
func (mr *MockComputeClientMockRecorder) DeleteFlavor(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFlavor", reflect.TypeOf((*MockComputeClient)(nil).DeleteFlavor), ctx, id)
}

// DeleteServer mocks base method.
func (m *MockComputeClient) DeleteServer(ctx context.Context, serverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", ctx, serverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServer indicates an expected call of DeleteServer.
func (mr *MockComputeClientMockRecorder) DeleteServer(ctx, serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockComputeClient)(nil).DeleteServer), ctx, serverID)
}

// GetFlavor mocks base method.
func (m *MockComputeClient) GetFlavor(ctx context.Context, id string) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlavor", ctx, id)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlavor indicates an expected call of GetFlavor.
func (mr *MockComputeClientMockRecorder) GetFlavor(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlavor", reflect.TypeOf((*MockComputeClient)(nil).GetFlavor), ctx, id)
}

// GetServer mocks base method.
func (m *MockComputeClient) GetServer(ctx context.Context, serverID string) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer", ctx, serverID)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServer indicates an expected call of GetServer.
func (mr *MockComputeClientMockRecorder) GetServer(ctx, serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockComputeClient)(nil).GetServer), ctx, serverID)
}

// ListFlavors mocks base method.
func (m *MockComputeClient) ListFlavors(ctx context.Context, listOpts flavors.ListOptsBuilder) iter.Seq2[*flavors.Flavor, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFlavors", ctx, listOpts)
	ret0, _ := ret[0].(iter.Seq2[*flavors.Flavor, error])
	return ret0
}

// ListFlavors indicates an expected call of ListFlavors.
func (mr *MockComputeClientMockRecorder) ListFlavors(ctx, listOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlavors", reflect.TypeOf((*MockComputeClient)(nil).ListFlavors), ctx, listOpts)
}

// ListServers mocks base method.
func (m *MockComputeClient) ListServers(ctx context.Context, listOpts servers.ListOptsBuilder) iter.Seq2[*servers.Server, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", ctx, listOpts)
	ret0, _ := ret[0].(iter.Seq2[*servers.Server, error])
	return ret0
}

// ListServers indicates an expected call of ListServers.
func (mr *MockComputeClientMockRecorder) ListServers(ctx, listOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockComputeClient)(nil).ListServers), ctx, listOpts)
}
