// Code generated by MockGen. DO NOT EDIT.
// Source: cache_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCacheRepository is a mock of CacheRepository interface.
type MockCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepositoryMockRecorder
}

// MockCacheRepositoryMockRecorder is the mock recorder for MockCacheRepository.
type MockCacheRepositoryMockRecorder struct {
	mock *MockCacheRepository
}

// NewMockCacheRepository creates a new mock instance.
func NewMockCacheRepository(ctrl *gomock.Controller) *MockCacheRepository {
	mock := &MockCacheRepository{ctrl: ctrl}
	mock.recorder = &MockCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheRepository) EXPECT() *MockCacheRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCacheRepository) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheRepositoryMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheRepository)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockCacheRepository) Get(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheRepositoryMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheRepository)(nil).Get), key)
}

// Set mocks base method.
func (m *MockCacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheRepositoryMockRecorder) Set(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheRepository)(nil).Set), key, value, expiration)
}