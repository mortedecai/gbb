// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mortedecai/go-burn-bits/gbb (interfaces: GBBClient)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_client.go -package=mocks github.com/mortedecai/go-burn-bits/gbb GBBClient
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGBBClient is a mock of GBBClient interface.
type MockGBBClient struct {
	ctrl     *gomock.Controller
	recorder *MockGBBClientMockRecorder
}

// MockGBBClientMockRecorder is the mock recorder for MockGBBClient.
type MockGBBClientMockRecorder struct {
	mock *MockGBBClient
}

// NewMockGBBClient creates a new mock instance.
func NewMockGBBClient(ctrl *gomock.Controller) *MockGBBClient {
	mock := &MockGBBClient{ctrl: ctrl}
	mock.recorder = &MockGBBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGBBClient) EXPECT() *MockGBBClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockGBBClient) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockGBBClientMockRecorder) Do(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockGBBClient)(nil).Do), arg0)
}
