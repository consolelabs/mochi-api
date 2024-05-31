package adapter

type ExecutionState string

const (
	ExecutionStatePending   ExecutionState = "QUERY_STATE_PENDING"
	ExecutionStateCompleted ExecutionState = "QUERY_STATE_COMPLETED"
)

type ExecutionResultMetadata struct {
	ColumnNames []string `json:"column_names"`
	ColumnTypes []string `json:"column_types"`
}

// ExecuteQueryResponse represents the response of the execute query endpoint
type ExecuteQueryResponse struct {
	ExecutionId string `json:"execution_id"`
	State       string `json:"state"`
}

// GetExecutionResultResponse represents the response of the get execution result endpoint
type GetExecutionResultResponse struct {
	ExecutionId         string           `json:"execution_id"`
	QueryId             int64            `json:"query_id"`
	IsExecutionFinished bool             `json:"is_execution_finished"`
	State               string           `json:"state"`
	Result              *ExecutionResult `json:"result"`
	NextOffset          *int64           `json:"next_offset"`
}

type ExecutionResult struct {
	Rows     []Row                   `json:"rows"`
	Metadata ExecutionResultMetadata `json:"metadata"`
}

type Row map[string]interface{}

// GetExecutionStatusResponse represents the response of the get execution status endpoint
type GetExecutionStatusResponse struct {
	ExecutionId         string                   `json:"execution_id"`
	QueryId             int64                    `json:"query_id"`
	IsExecutionFinished bool                     `json:"is_execution_finished"`
	State               string                   `json:"state"`
	ResultMetadata      *ExecutionResultMetadata `json:"result_metadata"`
}
