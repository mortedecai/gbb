// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mortedecai/gbb/commands (interfaces: Flagger)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_flagger.go -package=mocks github.com/mortedecai/gbb/commands Flagger
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cobra "github.com/spf13/cobra"
	gomock "go.uber.org/mock/gomock"
)

// MockFlagger is a mock of Flagger interface.
type MockFlagger struct {
	ctrl     *gomock.Controller
	recorder *MockFlaggerMockRecorder
}

// MockFlaggerMockRecorder is the mock recorder for MockFlagger.
type MockFlaggerMockRecorder struct {
	mock *MockFlagger
}

// NewMockFlagger creates a new mock instance.
func NewMockFlagger(ctrl *gomock.Controller) *MockFlagger {
	mock := &MockFlagger{ctrl: ctrl}
	mock.recorder = &MockFlaggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlagger) EXPECT() *MockFlaggerMockRecorder {
	return m.recorder
}

// GetInt mocks base method.
func (m *MockFlagger) GetInt(arg0 *cobra.Command, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt.
func (mr *MockFlaggerMockRecorder) GetInt(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockFlagger)(nil).GetInt), arg0, arg1)
}

// GetString mocks base method.
func (m *MockFlagger) GetString(arg0 *cobra.Command, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString.
func (mr *MockFlaggerMockRecorder) GetString(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockFlagger)(nil).GetString), arg0, arg1)
}
