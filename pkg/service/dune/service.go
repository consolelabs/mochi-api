package dune

import "github.com/defipod/mochi/pkg/service/dune/adapter"

type Service interface {
	ExecuteQuery(queryId int64, param map[string]any) (res *adapter.ExecuteQueryResponse, err error)
	GetExecutionResult(executionId string, limit, offset int64) (res *adapter.GetExecutionResultResponse, err error)
	GetExecutionStatus(executionId string) (res *adapter.GetExecutionResultResponse, err error)
}
