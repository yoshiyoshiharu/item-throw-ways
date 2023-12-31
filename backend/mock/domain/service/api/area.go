// Code generated by MockGen. DO NOT EDIT.
// Source: domain/service/api/area.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

// MockAreaService is a mock of AreaService interface.
type MockAreaService struct {
	ctrl     *gomock.Controller
	recorder *MockAreaServiceMockRecorder
}

// MockAreaServiceMockRecorder is the mock recorder for MockAreaService.
type MockAreaServiceMockRecorder struct {
	mock *MockAreaService
}

// NewMockAreaService creates a new mock instance.
func NewMockAreaService(ctrl *gomock.Controller) *MockAreaService {
	mock := &MockAreaService{ctrl: ctrl}
	mock.recorder = &MockAreaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAreaService) EXPECT() *MockAreaServiceMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockAreaService) FindAll() []*entity.Area {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Area)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAreaServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAreaService)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockAreaService) FindById(arg0 int) (*entity.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*entity.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAreaServiceMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAreaService)(nil).FindById), arg0)
}
