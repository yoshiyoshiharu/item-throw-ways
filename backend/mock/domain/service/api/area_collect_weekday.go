// Code generated by MockGen. DO NOT EDIT.
// Source: domain/service/api/area_collect_weekday.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

// MockAreaCollectWeekdayService is a mock of AreaCollectWeekdayService interface.
type MockAreaCollectWeekdayService struct {
	ctrl     *gomock.Controller
	recorder *MockAreaCollectWeekdayServiceMockRecorder
}

// MockAreaCollectWeekdayServiceMockRecorder is the mock recorder for MockAreaCollectWeekdayService.
type MockAreaCollectWeekdayServiceMockRecorder struct {
	mock *MockAreaCollectWeekdayService
}

// NewMockAreaCollectWeekdayService creates a new mock instance.
func NewMockAreaCollectWeekdayService(ctrl *gomock.Controller) *MockAreaCollectWeekdayService {
	mock := &MockAreaCollectWeekdayService{ctrl: ctrl}
	mock.recorder = &MockAreaCollectWeekdayServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaCollectWeekdayService) EXPECT() *MockAreaCollectWeekdayServiceMockRecorder {
	return m.recorder
}

// ConvertByAreaWithAroundMonths mocks base method.
func (m *MockAreaCollectWeekdayService) ConvertByAreaWithAroundMonths(arg0, arg1 int, arg2 time.Month) []*entity.AreaCollectDate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertByAreaWithAroundMonths", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entity.AreaCollectDate)
	return ret0
}

// ConvertByAreaWithAroundMonths indicates an expected call of ConvertByAreaWithAroundMonths.
func (mr *MockAreaCollectWeekdayServiceMockRecorder) ConvertByAreaWithAroundMonths(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertByAreaWithAroundMonths", reflect.TypeOf((*MockAreaCollectWeekdayService)(nil).ConvertByAreaWithAroundMonths), arg0, arg1, arg2)
}
