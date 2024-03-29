package response

import "time"

type ProductMetadataCopy struct {
	Id          int64       `json:"id"`
	Type        string      `json:"type"`
	Description interface{} `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
