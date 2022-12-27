package offchain_tip_bot_contract

type ListQuery struct {
	ChainID        string
	IsEVM          *bool
	SupportDeposit *bool
}
