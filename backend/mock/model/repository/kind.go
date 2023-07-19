// Code generated by MockGen. DO NOT EDIT.
// Source: model/repository/kind.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/yoshiyoshiharu/item-throw-ways/model/entity"
)

// MockKindRepository is a mock of KindRepository interface.
type MockKindRepository struct {
	ctrl     *gomock.Controller
	recorder *MockKindRepositoryMockRecorder
}

// MockKindRepositoryMockRecorder is the mock recorder for MockKindRepository.
type MockKindRepositoryMockRecorder struct {
	mock *MockKindRepository
}

// NewMockKindRepository creates a new mock instance.
func NewMockKindRepository(ctrl *gomock.Controller) *MockKindRepository {
	mock := &MockKindRepository{ctrl: ctrl}
	mock.recorder = &MockKindRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKindRepository) EXPECT() *MockKindRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockKindRepository) FindAll() []*entity.Kind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Kind)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockKindRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockKindRepository)(nil).FindAll))
}