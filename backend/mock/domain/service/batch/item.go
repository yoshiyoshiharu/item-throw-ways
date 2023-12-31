// Code generated by MockGen. DO NOT EDIT.
// Source: domain/service/batch/item.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockItemBatchService is a mock of ItemBatchService interface.
type MockItemBatchService struct {
	ctrl     *gomock.Controller
	recorder *MockItemBatchServiceMockRecorder
}

// MockItemBatchServiceMockRecorder is the mock recorder for MockItemBatchService.
type MockItemBatchServiceMockRecorder struct {
	mock *MockItemBatchService
}

// NewMockItemBatchService creates a new mock instance.
func NewMockItemBatchService(ctrl *gomock.Controller) *MockItemBatchService {
	mock := &MockItemBatchService{ctrl: ctrl}
	mock.recorder = &MockItemBatchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemBatchService) EXPECT() *MockItemBatchServiceMockRecorder {
	return m.recorder
}

// UpdateAll mocks base method.
func (m *MockItemBatchService) UpdateAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockItemBatchServiceMockRecorder) UpdateAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockItemBatchService)(nil).UpdateAll))
}
