// Code generated by MockGen. DO NOT EDIT.
// Source: domain/service/batch/area_collect_weekday.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAreaCollectWeekdayBatchService is a mock of AreaCollectWeekdayBatchService interface.
type MockAreaCollectWeekdayBatchService struct {
	ctrl     *gomock.Controller
	recorder *MockAreaCollectWeekdayBatchServiceMockRecorder
}

// MockAreaCollectWeekdayBatchServiceMockRecorder is the mock recorder for MockAreaCollectWeekdayBatchService.
type MockAreaCollectWeekdayBatchServiceMockRecorder struct {
	mock *MockAreaCollectWeekdayBatchService
}

// NewMockAreaCollectWeekdayBatchService creates a new mock instance.
func NewMockAreaCollectWeekdayBatchService(ctrl *gomock.Controller) *MockAreaCollectWeekdayBatchService {
	mock := &MockAreaCollectWeekdayBatchService{ctrl: ctrl}
	mock.recorder = &MockAreaCollectWeekdayBatchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaCollectWeekdayBatchService) EXPECT() *MockAreaCollectWeekdayBatchServiceMockRecorder {
	return m.recorder
}

// UpdateAll mocks base method.
func (m *MockAreaCollectWeekdayBatchService) UpdateAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAreaCollectWeekdayBatchServiceMockRecorder) UpdateAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAreaCollectWeekdayBatchService)(nil).UpdateAll))
}