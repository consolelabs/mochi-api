// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/guild_config_token/store.go

// Package mock_guild_config_token is a generated GoMock package.
package mock_guild_config_token

import (
	reflect "reflect"

	model "github.com/defipod/mochi/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetByGuildID mocks base method.
func (m *MockStore) GetByGuildID(guildID string) ([]model.GuildConfigToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGuildID", guildID)
	ret0, _ := ret[0].([]model.GuildConfigToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGuildID indicates an expected call of GetByGuildID.
func (mr *MockStoreMockRecorder) GetByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGuildID", reflect.TypeOf((*MockStore)(nil).GetByGuildID), guildID)
}

// UpsertMany mocks base method.
func (m *MockStore) UpsertMany(configs []model.GuildConfigToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", configs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockStoreMockRecorder) UpsertMany(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockStore)(nil).UpsertMany), configs)
}
