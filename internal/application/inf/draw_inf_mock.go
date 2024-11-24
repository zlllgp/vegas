// Code generated by MockGen. DO NOT EDIT.
// Source: draw_inf.go

// Package inf is a generated GoMock package.
package inf

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/zlllgp/vegas/internal/domain/entity"
)

// MockDrawService is a mock of DrawService interface.
type MockDrawService struct {
	ctrl     *gomock.Controller
	recorder *MockDrawServiceMockRecorder
}

// MockDrawServiceMockRecorder is the mock recorder for MockDrawService.
type MockDrawServiceMockRecorder struct {
	mock *MockDrawService
}

// NewMockDrawService creates a new mock instance.
func NewMockDrawService(ctrl *gomock.Controller) *MockDrawService {
	mock := &MockDrawService{ctrl: ctrl}
	mock.recorder = &MockDrawServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDrawService) EXPECT() *MockDrawServiceMockRecorder {
	return m.recorder
}

// Draw mocks base method.
func (m *MockDrawService) Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Draw", ctx, userId, eh)
	ret0, _ := ret[0].(*entity.DrawResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Draw indicates an expected call of Draw.
func (mr *MockDrawServiceMockRecorder) Draw(ctx, userId, eh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockDrawService)(nil).Draw), ctx, userId, eh)
}