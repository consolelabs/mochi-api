// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/quest/store.go

// Package mock_quest is a generated GoMock package.
package mock_quest

import (
	model "github.com/defipod/mochi/pkg/model"
	quest "github.com/defipod/mochi/pkg/repo/quest"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockStore) List(q quest.ListQuery) ([]model.Quest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", q)
	ret0, _ := ret[0].([]model.Quest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), q)
}

// GetAvailableRoutines mocks base method
func (m *MockStore) GetAvailableRoutines() ([]model.QuestRoutine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableRoutines")
	ret0, _ := ret[0].([]model.QuestRoutine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableRoutines indicates an expected call of GetAvailableRoutines
func (mr *MockStoreMockRecorder) GetAvailableRoutines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableRoutines", reflect.TypeOf((*MockStore)(nil).GetAvailableRoutines))
}