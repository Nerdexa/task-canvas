// Code generated by MockGen. DO NOT EDIT.
// Source: port/todo_id_port.go
//
// Generated by this command:
//
//	mockgen -source=port/todo_id_port.go -destination=mock/port/todo_id_port.go
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"
	domain "task-canvas/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockTodoIdPort is a mock of TodoIdPort interface.
type MockTodoIdPort struct {
	ctrl     *gomock.Controller
	recorder *MockTodoIdPortMockRecorder
}

// MockTodoIdPortMockRecorder is the mock recorder for MockTodoIdPort.
type MockTodoIdPortMockRecorder struct {
	mock *MockTodoIdPort
}

// NewMockTodoIdPort creates a new mock instance.
func NewMockTodoIdPort(ctrl *gomock.Controller) *MockTodoIdPort {
	mock := &MockTodoIdPort{ctrl: ctrl}
	mock.recorder = &MockTodoIdPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoIdPort) EXPECT() *MockTodoIdPortMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTodoIdPort) Generate(ctx context.Context) (domain.TodoId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx)
	ret0, _ := ret[0].(domain.TodoId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTodoIdPortMockRecorder) Generate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTodoIdPort)(nil).Generate), ctx)
}
