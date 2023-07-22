// Code generated by MockGen. DO NOT EDIT.
// Source: model/repository/item.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// DeleteAndInsertAll mocks base method.
func (m *MockItemRepository) DeleteAndInsertAll(arg0 []entity.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndInsertAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndInsertAll indicates an expected call of DeleteAndInsertAll.
func (mr *MockItemRepositoryMockRecorder) DeleteAndInsertAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndInsertAll", reflect.TypeOf((*MockItemRepository)(nil).DeleteAndInsertAll), arg0)
}

// FindAll mocks base method.
func (m *MockItemRepository) FindAll() []*entity.Item {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Item)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockItemRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockItemRepository)(nil).FindAll))
}
