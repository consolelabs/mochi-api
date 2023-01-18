package request

type CreateEnvelop struct {
	UserID  string `json:"user_id" binding:"required"`
	Command string `json:"command" binding:"required" `
}
