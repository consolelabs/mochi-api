// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repo/token/store.go

// Package mock_token is a generated GoMock package.
package mock_token

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

// CreateOne mocks base method.
func (m *MockStore) CreateOne(token *model.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOne indicates an expected call of CreateOne.
func (mr *MockStoreMockRecorder) CreateOne(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockStore)(nil).CreateOne), token)
}

// GetAll mocks base method.
func (m *MockStore) GetAll() ([]model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStore)(nil).GetAll))
}

// GetAllSupported mocks base method.
func (m *MockStore) GetAllSupported() ([]model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSupported")
	ret0, _ := ret[0].([]model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSupported indicates an expected call of GetAllSupported.
func (mr *MockStoreMockRecorder) GetAllSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSupported", reflect.TypeOf((*MockStore)(nil).GetAllSupported))
}

// GetByAddress mocks base method.
func (m *MockStore) GetByAddress(address string, chainID int) (*model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddress", address, chainID)
	ret0, _ := ret[0].(*model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddress indicates an expected call of GetByAddress.
func (mr *MockStoreMockRecorder) GetByAddress(address, chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddress", reflect.TypeOf((*MockStore)(nil).GetByAddress), address, chainID)
}

// GetBySymbol mocks base method.
func (m *MockStore) GetBySymbol(symbol string, botSupported bool) (model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySymbol", symbol, botSupported)
	ret0, _ := ret[0].(model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySymbol indicates an expected call of GetBySymbol.
func (mr *MockStoreMockRecorder) GetBySymbol(symbol, botSupported interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySymbol", reflect.TypeOf((*MockStore)(nil).GetBySymbol), symbol, botSupported)
}

// GetDefaultTokenByGuildID mocks base method.
func (m *MockStore) GetDefaultTokenByGuildID(guildID string) (model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultTokenByGuildID", guildID)
	ret0, _ := ret[0].(model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultTokenByGuildID indicates an expected call of GetDefaultTokenByGuildID.
func (mr *MockStoreMockRecorder) GetDefaultTokenByGuildID(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultTokenByGuildID", reflect.TypeOf((*MockStore)(nil).GetDefaultTokenByGuildID), guildID)
}

// GetDefaultTokens mocks base method.
func (m *MockStore) GetDefaultTokens() ([]model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultTokens")
	ret0, _ := ret[0].([]model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultTokens indicates an expected call of GetDefaultTokens.
func (mr *MockStoreMockRecorder) GetDefaultTokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultTokens", reflect.TypeOf((*MockStore)(nil).GetDefaultTokens))
}

// GetOneBySymbol mocks base method.
func (m *MockStore) GetOneBySymbol(symbol string) (*model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneBySymbol", symbol)
	ret0, _ := ret[0].(*model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneBySymbol indicates an expected call of GetOneBySymbol.
func (mr *MockStoreMockRecorder) GetOneBySymbol(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneBySymbol", reflect.TypeOf((*MockStore)(nil).GetOneBySymbol), symbol)
}

// GetSupportedTokenByGuildId mocks base method.
func (m *MockStore) GetSupportedTokenByGuildId(guildID string) ([]model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedTokenByGuildId", guildID)
	ret0, _ := ret[0].([]model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedTokenByGuildId indicates an expected call of GetSupportedTokenByGuildId.
func (mr *MockStoreMockRecorder) GetSupportedTokenByGuildId(guildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedTokenByGuildId", reflect.TypeOf((*MockStore)(nil).GetSupportedTokenByGuildId), guildID)
}

// UpsertOne mocks base method.
func (m *MockStore) UpsertOne(token model.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertOne", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertOne indicates an expected call of UpsertOne.
func (mr *MockStoreMockRecorder) UpsertOne(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertOne", reflect.TypeOf((*MockStore)(nil).UpsertOne), token)
}
