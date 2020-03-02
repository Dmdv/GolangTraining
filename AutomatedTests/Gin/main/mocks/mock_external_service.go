// Code generated by MockGen. DO NOT EDIT.
// Source: external_service/external_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExternalService is a mock of ExternalService interface
type MockExternalService struct {
	ctrl     *gomock.Controller
	recorder *MockExternalServiceMockRecorder
}

// MockExternalServiceMockRecorder is the mock recorder for MockExternalService
type MockExternalServiceMockRecorder struct {
	mock *MockExternalService
}

// NewMockExternalService creates a new mock instance
func NewMockExternalService(ctrl *gomock.Controller) *MockExternalService {
	mock := &MockExternalService{ctrl: ctrl}
	mock.recorder = &MockExternalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalService) EXPECT() *MockExternalServiceMockRecorder {
	return m.recorder
}

// ReadPage mocks base method
func (m *MockExternalService) ReadPage(url string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPage", url)
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadPage indicates an expected call of ReadPage
func (mr *MockExternalServiceMockRecorder) ReadPage(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPage", reflect.TypeOf((*MockExternalService)(nil).ReadPage), url)
}
