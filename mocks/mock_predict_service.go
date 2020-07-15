// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fake-eta-task/internal/predict (interfaces: PredictService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "github.com/fake-eta-task/internal/predict/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPredictService is a mock of PredictService interface
type MockPredictService struct {
	ctrl     *gomock.Controller
	recorder *MockPredictServiceMockRecorder
}

// MockPredictServiceMockRecorder is the mock recorder for MockPredictService
type MockPredictServiceMockRecorder struct {
	mock *MockPredictService
}

// NewMockPredictService creates a new mock instance
func NewMockPredictService(ctrl *gomock.Controller) *MockPredictService {
	mock := &MockPredictService{ctrl: ctrl}
	mock.recorder = &MockPredictServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPredictService) EXPECT() *MockPredictServiceMockRecorder {
	return m.recorder
}

// PredictWaitingTimes mocks base method
func (m *MockPredictService) PredictWaitingTimes(arg0 context.Context, arg1 []models.Position, arg2 models.Position) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PredictWaitingTimes", arg0, arg1, arg2)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PredictWaitingTimes indicates an expected call of PredictWaitingTimes
func (mr *MockPredictServiceMockRecorder) PredictWaitingTimes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PredictWaitingTimes", reflect.TypeOf((*MockPredictService)(nil).PredictWaitingTimes), arg0, arg1, arg2)
}