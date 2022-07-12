// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/indexer/service.go

// Package mock_indexer is a generated GoMock package.
package mock_indexer

import (
	reflect "reflect"

	response "github.com/defipod/mochi/pkg/response"
	indexer "github.com/defipod/mochi/pkg/service/indexer"
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

// CreateERC721Contract mocks base method.
func (m *MockService) CreateERC721Contract(arg0 indexer.CreateERC721ContractRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateERC721Contract", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateERC721Contract indicates an expected call of CreateERC721Contract.
func (mr *MockServiceMockRecorder) CreateERC721Contract(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateERC721Contract", reflect.TypeOf((*MockService)(nil).CreateERC721Contract), arg0)
}

// GetAttributeIcon mocks base method.
func (m *MockService) GetAttributeIcon() (*response.AttributeIconResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttributeIcon")
	ret0, _ := ret[0].(*response.AttributeIconResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttributeIcon indicates an expected call of GetAttributeIcon.
func (mr *MockServiceMockRecorder) GetAttributeIcon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributeIcon", reflect.TypeOf((*MockService)(nil).GetAttributeIcon))
}

// GetNFTCollectionTickers mocks base method.
func (m *MockService) GetNFTCollectionTickers(address, rawQuery string) (*response.IndexerNFTCollectionTickersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTCollectionTickers", address, rawQuery)
	ret0, _ := ret[0].(*response.IndexerNFTCollectionTickersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTCollectionTickers indicates an expected call of GetNFTCollectionTickers.
func (mr *MockServiceMockRecorder) GetNFTCollectionTickers(address, rawQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTCollectionTickers", reflect.TypeOf((*MockService)(nil).GetNFTCollectionTickers), address, rawQuery)
}

// GetNFTCollections mocks base method.
func (m *MockService) GetNFTCollections(query string) (*response.IndexerGetNFTCollectionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTCollections", query)
	ret0, _ := ret[0].(*response.IndexerGetNFTCollectionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTCollections indicates an expected call of GetNFTCollections.
func (mr *MockServiceMockRecorder) GetNFTCollections(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTCollections", reflect.TypeOf((*MockService)(nil).GetNFTCollections), query)
}

// GetNFTContract mocks base method.
func (m *MockService) GetNFTContract(addr string) (*response.IndexerContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTContract", addr)
	ret0, _ := ret[0].(*response.IndexerContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTContract indicates an expected call of GetNFTContract.
func (mr *MockServiceMockRecorder) GetNFTContract(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTContract", reflect.TypeOf((*MockService)(nil).GetNFTContract), addr)
}

// GetNFTDetail mocks base method.
func (m *MockService) GetNFTDetail(collectionAddress, tokenID string) (*response.IndexerNFTToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTDetail", collectionAddress, tokenID)
	ret0, _ := ret[0].(*response.IndexerNFTToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTDetail indicates an expected call of GetNFTDetail.
func (mr *MockServiceMockRecorder) GetNFTDetail(collectionAddress, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTDetail", reflect.TypeOf((*MockService)(nil).GetNFTDetail), collectionAddress, tokenID)
}

// GetNFTTokens mocks base method.
func (m *MockService) GetNFTTokens(address, query string) (*response.IndexerGetNFTTokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTTokens", address, query)
	ret0, _ := ret[0].(*response.IndexerGetNFTTokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTTokens indicates an expected call of GetNFTTokens.
func (mr *MockServiceMockRecorder) GetNFTTokens(address, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTTokens", reflect.TypeOf((*MockService)(nil).GetNFTTokens), address, query)
}

// GetNFTTradingVolume mocks base method.
func (m *MockService) GetNFTTradingVolume() ([]response.NFTTradingVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTTradingVolume")
	ret0, _ := ret[0].([]response.NFTTradingVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNFTTradingVolume indicates an expected call of GetNFTTradingVolume.
func (mr *MockServiceMockRecorder) GetNFTTradingVolume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTTradingVolume", reflect.TypeOf((*MockService)(nil).GetNFTTradingVolume))
}

// GetNftSales mocks base method.
func (m *MockService) GetNftSales(addr, platform string) (*response.NftSalesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftSales", addr, platform)
	ret0, _ := ret[0].(*response.NftSalesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftSales indicates an expected call of GetNftSales.
func (mr *MockServiceMockRecorder) GetNftSales(addr, platform interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftSales", reflect.TypeOf((*MockService)(nil).GetNftSales), addr, platform)
}
