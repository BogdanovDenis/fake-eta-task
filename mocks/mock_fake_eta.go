// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fake-eta-task/restapi (interfaces: FakeEta)

// Package mocks is a generated GoMock package.
package mocks

import (
	operations "github.com/fake-eta-task/restapi/operations"
	middleware "github.com/go-openapi/runtime/middleware"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFakeEta is a mock of FakeEta interface
type MockFakeEta struct {
	ctrl     *gomock.Controller
	recorder *MockFakeEtaMockRecorder
}

// MockFakeEtaMockRecorder is the mock recorder for MockFakeEta
type MockFakeEtaMockRecorder struct {
	mock *MockFakeEta
}

// NewMockFakeEta creates a new mock instance
func NewMockFakeEta(ctrl *gomock.Controller) *MockFakeEta {
	mock := &MockFakeEta{ctrl: ctrl}
	mock.recorder = &MockFakeEtaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFakeEta) EXPECT() *MockFakeEtaMockRecorder {
	return m.recorder
}

// CalculateWaitingTimeImpl mocks base method
func (m *MockFakeEta) CalculateWaitingTimeImpl(arg0 operations.CalculateWaitingTimeParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateWaitingTimeImpl", arg0)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// CalculateWaitingTimeImpl indicates an expected call of CalculateWaitingTimeImpl
func (mr *MockFakeEtaMockRecorder) CalculateWaitingTimeImpl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateWaitingTimeImpl", reflect.TypeOf((*MockFakeEta)(nil).CalculateWaitingTimeImpl), arg0)
}
