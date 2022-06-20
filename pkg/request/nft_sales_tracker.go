package request

type NFTSalesTrackerRequest struct{
	ContractAddress string `json:"contract_address"`
	Platform	string `json:"platform"`
	SalesConfigID string `json:"sales_config_id"`
} 