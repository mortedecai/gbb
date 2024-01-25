// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mortedecai/gbb/client (interfaces: DownloadOption)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_download_option.go -package=mocks github.com/mortedecai/gbb/client DownloadOption
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDownloadOption is a mock of DownloadOption interface.
type MockDownloadOption struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadOptionMockRecorder
}

// MockDownloadOptionMockRecorder is the mock recorder for MockDownloadOption.
type MockDownloadOptionMockRecorder struct {
	mock *MockDownloadOption
}

// NewMockDownloadOption creates a new mock instance.
func NewMockDownloadOption(ctrl *gomock.Controller) *MockDownloadOption {
	mock := &MockDownloadOption{ctrl: ctrl}
	mock.recorder = &MockDownloadOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadOption) EXPECT() *MockDownloadOptionMockRecorder {
	return m.recorder
}

// AddAuth mocks base method.
func (m *MockDownloadOption) AddAuth(arg0 *http.Request) *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuth", arg0)
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// AddAuth indicates an expected call of AddAuth.
func (mr *MockDownloadOptionMockRecorder) AddAuth(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuth", reflect.TypeOf((*MockDownloadOption)(nil).AddAuth), arg0)
}

// AuthToken mocks base method.
func (m *MockDownloadOption) AuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthToken indicates an expected call of AuthToken.
func (mr *MockDownloadOptionMockRecorder) AuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthToken", reflect.TypeOf((*MockDownloadOption)(nil).AuthToken))
}

// Destination mocks base method.
func (m *MockDownloadOption) Destination() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destination")
	ret0, _ := ret[0].(string)
	return ret0
}

// Destination indicates an expected call of Destination.
func (mr *MockDownloadOptionMockRecorder) Destination() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destination", reflect.TypeOf((*MockDownloadOption)(nil).Destination))
}

// Host mocks base method.
func (m *MockDownloadOption) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockDownloadOptionMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockDownloadOption)(nil).Host))
}

// Port mocks base method.
func (m *MockDownloadOption) Port() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Port")
	ret0, _ := ret[0].(int)
	return ret0
}

// Port indicates an expected call of Port.
func (mr *MockDownloadOptionMockRecorder) Port() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockDownloadOption)(nil).Port))
}

// Valid mocks base method.
func (m *MockDownloadOption) Valid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockDownloadOptionMockRecorder) Valid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockDownloadOption)(nil).Valid))
}
