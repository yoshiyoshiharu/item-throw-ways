// Code generated by MockGen. DO NOT EDIT.
// Source: model/repository/area_collect_weekday.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

// MockAreaCollectWeekdayRepository is a mock of AreaCollectWeekdayRepository interface.
type MockAreaCollectWeekdayRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAreaCollectWeekdayRepositoryMockRecorder
}

// MockAreaCollectWeekdayRepositoryMockRecorder is the mock recorder for MockAreaCollectWeekdayRepository.
type MockAreaCollectWeekdayRepositoryMockRecorder struct {
	mock *MockAreaCollectWeekdayRepository
}

// NewMockAreaCollectWeekdayRepository creates a new mock instance.
func NewMockAreaCollectWeekdayRepository(ctrl *gomock.Controller) *MockAreaCollectWeekdayRepository {
	mock := &MockAreaCollectWeekdayRepository{ctrl: ctrl}
	mock.recorder = &MockAreaCollectWeekdayRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaCollectWeekdayRepository) EXPECT() *MockAreaCollectWeekdayRepositoryMockRecorder {
	return m.recorder
}

// FindByAreaId mocks base method.
func (m *MockAreaCollectWeekdayRepository) FindByAreaId(arg0 int) []*entity.AreaCollectWeekday {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAreaId", arg0)
	ret0, _ := ret[0].([]*entity.AreaCollectWeekday)
	return ret0
}

// FindByAreaId indicates an expected call of FindByAreaId.
func (mr *MockAreaCollectWeekdayRepositoryMockRecorder) FindByAreaId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAreaId", reflect.TypeOf((*MockAreaCollectWeekdayRepository)(nil).FindByAreaId), arg0)
}
