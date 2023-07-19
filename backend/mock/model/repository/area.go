// Code generated by MockGen. DO NOT EDIT.
// Source: model/repository/area.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

// MockAreaRepository is a mock of AreaRepository interface.
type MockAreaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAreaRepositoryMockRecorder
}

// MockAreaRepositoryMockRecorder is the mock recorder for MockAreaRepository.
type MockAreaRepositoryMockRecorder struct {
	mock *MockAreaRepository
}

// NewMockAreaRepository creates a new mock instance.
func NewMockAreaRepository(ctrl *gomock.Controller) *MockAreaRepository {
	mock := &MockAreaRepository{ctrl: ctrl}
	mock.recorder = &MockAreaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaRepository) EXPECT() *MockAreaRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockAreaRepository) FindAll() []*entity.Area {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Area)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAreaRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAreaRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockAreaRepository) FindById(arg0 int) (*entity.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*entity.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAreaRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAreaRepository)(nil).FindById), arg0)
}
