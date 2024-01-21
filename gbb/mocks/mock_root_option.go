// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mortedecai/gbb/gbb (interfaces: CommandOption)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_root_option.go -package=mocks github.com/mortedecai/gbb/gbb CommandOption
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCommandOption is a mock of CommandOption interface.
type MockCommandOption struct {
	ctrl     *gomock.Controller
	recorder *MockCommandOptionMockRecorder
}

// MockCommandOptionMockRecorder is the mock recorder for MockCommandOption.
type MockCommandOptionMockRecorder struct {
	mock *MockCommandOption
}

// NewMockCommandOption creates a new mock instance.
func NewMockCommandOption(ctrl *gomock.Controller) *MockCommandOption {
	mock := &MockCommandOption{ctrl: ctrl}
	mock.recorder = &MockCommandOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandOption) EXPECT() *MockCommandOptionMockRecorder {
	return m.recorder
}

// AddAuth mocks base method.
func (m *MockCommandOption) AddAuth(arg0 *http.Request) *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuth", arg0)
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// AddAuth indicates an expected call of AddAuth.
func (mr *MockCommandOptionMockRecorder) AddAuth(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuth", reflect.TypeOf((*MockCommandOption)(nil).AddAuth), arg0)
}

// AuthToken mocks base method.
func (m *MockCommandOption) AuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthToken indicates an expected call of AuthToken.
func (mr *MockCommandOptionMockRecorder) AuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthToken", reflect.TypeOf((*MockCommandOption)(nil).AuthToken))
}

// Host mocks base method.
func (m *MockCommandOption) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockCommandOptionMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockCommandOption)(nil).Host))
}

// Port mocks base method.
func (m *MockCommandOption) Port() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Port")
	ret0, _ := ret[0].(int)
	return ret0
}

// Port indicates an expected call of Port.
func (mr *MockCommandOptionMockRecorder) Port() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockCommandOption)(nil).Port))
}

// Valid mocks base method.
func (m *MockCommandOption) Valid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockCommandOptionMockRecorder) Valid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockCommandOption)(nil).Valid))
}
