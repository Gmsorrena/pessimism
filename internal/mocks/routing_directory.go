// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/alert (interfaces: RoutingDirectory)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	client "github.com/base-org/pessimism/internal/client"
	core "github.com/base-org/pessimism/internal/core"
	gomock "github.com/golang/mock/gomock"
)

// MockRoutingDirectory is a mock of RoutingDirectory interface.
type MockRoutingDirectory struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingDirectoryMockRecorder
}

// MockRoutingDirectoryMockRecorder is the mock recorder for MockRoutingDirectory.
type MockRoutingDirectoryMockRecorder struct {
	mock *MockRoutingDirectory
}

// NewMockRoutingDirectory creates a new mock instance.
func NewMockRoutingDirectory(ctrl *gomock.Controller) *MockRoutingDirectory {
	mock := &MockRoutingDirectory{ctrl: ctrl}
	mock.recorder = &MockRoutingDirectoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutingDirectory) EXPECT() *MockRoutingDirectoryMockRecorder {
	return m.recorder
}

// GetPagerDutyClients mocks base method.
func (m *MockRoutingDirectory) GetPagerDutyClients(arg0 core.Severity) []client.PagerDutyClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPagerDutyClients", arg0)
	ret0, _ := ret[0].([]client.PagerDutyClient)
	return ret0
}

// GetPagerDutyClients indicates an expected call of GetPagerDutyClients.
func (mr *MockRoutingDirectoryMockRecorder) GetPagerDutyClients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPagerDutyClients", reflect.TypeOf((*MockRoutingDirectory)(nil).GetPagerDutyClients), arg0)
}

// GetSlackClients mocks base method.
func (m *MockRoutingDirectory) GetSlackClients(arg0 core.Severity) []client.SlackClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlackClients", arg0)
	ret0, _ := ret[0].([]client.SlackClient)
	return ret0
}

// GetSlackClients indicates an expected call of GetSlackClients.
func (mr *MockRoutingDirectoryMockRecorder) GetSlackClients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlackClients", reflect.TypeOf((*MockRoutingDirectory)(nil).GetSlackClients), arg0)
}

// InitializeRouting mocks base method.
func (m *MockRoutingDirectory) InitializeRouting(arg0 *core.AlertRoutingParams) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitializeRouting", arg0)
}

// InitializeRouting indicates an expected call of InitializeRouting.
func (mr *MockRoutingDirectoryMockRecorder) InitializeRouting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeRouting", reflect.TypeOf((*MockRoutingDirectory)(nil).InitializeRouting), arg0)
}

// SetPagerDutyClients mocks base method.
func (m *MockRoutingDirectory) SetPagerDutyClients(arg0 []client.PagerDutyClient, arg1 core.Severity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPagerDutyClients", arg0, arg1)
}

// SetPagerDutyClients indicates an expected call of SetPagerDutyClients.
func (mr *MockRoutingDirectoryMockRecorder) SetPagerDutyClients(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPagerDutyClients", reflect.TypeOf((*MockRoutingDirectory)(nil).SetPagerDutyClients), arg0, arg1)
}

// SetSlackClients mocks base method.
func (m *MockRoutingDirectory) SetSlackClients(arg0 []client.SlackClient, arg1 core.Severity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSlackClients", arg0, arg1)
}

// SetSlackClients indicates an expected call of SetSlackClients.
func (mr *MockRoutingDirectoryMockRecorder) SetSlackClients(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSlackClients", reflect.TypeOf((*MockRoutingDirectory)(nil).SetSlackClients), arg0, arg1)
}