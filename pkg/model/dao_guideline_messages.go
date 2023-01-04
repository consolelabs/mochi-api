package model

type DaoGuidelineMessage struct {
	Id        int64                 `json:"id"`
	Authority ProposalAuthorityType `json:"authority"`
	Message   string                `json:"message"`
}

func (DaoGuidelineMessage) TableName() string {
	return "dao_guideline_messages"
}
