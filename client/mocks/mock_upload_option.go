// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mortedecai/gbb/client (interfaces: UploadOption)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_upload_option.go -package=mocks github.com/mortedecai/gbb/client UploadOption
//
// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	models "github.com/mortedecai/gbb/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUploadOption is a mock of UploadOption interface.
type MockUploadOption struct {
	ctrl     *gomock.Controller
	recorder *MockUploadOptionMockRecorder
}

// MockUploadOptionMockRecorder is the mock recorder for MockUploadOption.
type MockUploadOptionMockRecorder struct {
	mock *MockUploadOption
}

// NewMockUploadOption creates a new mock instance.
func NewMockUploadOption(ctrl *gomock.Controller) *MockUploadOption {
	mock := &MockUploadOption{ctrl: ctrl}
	mock.recorder = &MockUploadOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadOption) EXPECT() *MockUploadOptionMockRecorder {
	return m.recorder
}

// AddAuth mocks base method.
func (m *MockUploadOption) AddAuth(arg0 *http.Request) *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAuth", arg0)
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// AddAuth indicates an expected call of AddAuth.
func (mr *MockUploadOptionMockRecorder) AddAuth(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAuth", reflect.TypeOf((*MockUploadOption)(nil).AddAuth), arg0)
}

// AuthToken mocks base method.
func (m *MockUploadOption) AuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthToken indicates an expected call of AuthToken.
func (mr *MockUploadOptionMockRecorder) AuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthToken", reflect.TypeOf((*MockUploadOption)(nil).AuthToken))
}

// Host mocks base method.
func (m *MockUploadOption) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockUploadOptionMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockUploadOption)(nil).Host))
}

// Port mocks base method.
func (m *MockUploadOption) Port() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Port")
	ret0, _ := ret[0].(int)
	return ret0
}

// Port indicates an expected call of Port.
func (mr *MockUploadOptionMockRecorder) Port() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockUploadOption)(nil).Port))
}

// Server mocks base method.
func (m *MockUploadOption) Server() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(string)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockUploadOptionMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockUploadOption)(nil).Server))
}

// ToUpload mocks base method.
func (m *MockUploadOption) ToUpload() []models.GBBFileName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUpload")
	ret0, _ := ret[0].([]models.GBBFileName)
	return ret0
}

// ToUpload indicates an expected call of ToUpload.
func (mr *MockUploadOptionMockRecorder) ToUpload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUpload", reflect.TypeOf((*MockUploadOption)(nil).ToUpload))
}

// Valid mocks base method.
func (m *MockUploadOption) Valid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockUploadOptionMockRecorder) Valid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockUploadOption)(nil).Valid))
}