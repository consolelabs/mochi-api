package response

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Transactions struct {
	ID          uuid.UUID       `json:"id" swaggertype:"string"`
	SenderID    string          `json:"sender_id"`
	ReceiverID  string          `json:"receiver_id"`
	GuildID     string          `json:"guild_id"`
	LogID       string          `json:"log_id"`
	Status      string          `json:"status"`
	Amount      float64         `json:"amount"`
	Token       string          `json:"token"`
	Action      string          `json:"action"`
	ServiceFee  sql.NullFloat64 `json:"service_fee"`
	FeeAmount   sql.NullFloat64 `json:"fee_amount"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	FullCommand string          `json:"full_command"`
}

type TransactionsResponse struct {
	Data []TransactionsResponse `json:"data"`
}
