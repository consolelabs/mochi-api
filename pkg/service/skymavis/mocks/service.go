// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/skymavis/service.go

// Package mock_skymavis is a generated GoMock package.
package mock_skymavis

import (
	reflect "reflect"

	response "github.com/defipod/mochi/pkg/response"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetAddressFarming mocks base method.
func (m *MockService) GetAddressFarming(address string) (*response.WalletFarmingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressFarming", address)
	ret0, _ := ret[0].(*response.WalletFarmingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressFarming indicates an expected call of GetAddressFarming.
func (mr *MockServiceMockRecorder) GetAddressFarming(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressFarming", reflect.TypeOf((*MockService)(nil).GetAddressFarming), address)
}

// GetInternalTxnsByHash mocks base method.
func (m *MockService) GetInternalTxnsByHash(hash string) (*response.SkymavisTransactionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternalTxnsByHash", hash)
	ret0, _ := ret[0].(*response.SkymavisTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInternalTxnsByHash indicates an expected call of GetInternalTxnsByHash.
func (mr *MockServiceMockRecorder) GetInternalTxnsByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternalTxnsByHash", reflect.TypeOf((*MockService)(nil).GetInternalTxnsByHash), hash)
}

// GetOwnedNfts mocks base method.
func (m *MockService) GetOwnedNfts(address string) (*response.AxieMarketNftResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnedNfts", address)
	ret0, _ := ret[0].(*response.AxieMarketNftResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwnedNfts indicates an expected call of GetOwnedNfts.
func (mr *MockServiceMockRecorder) GetOwnedNfts(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnedNfts", reflect.TypeOf((*MockService)(nil).GetOwnedNfts), address)
}
