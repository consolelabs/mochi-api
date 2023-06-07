package request

type GeneratePkPassRequest struct {
	Category string `form:"category" binding:"required"`
	Type     string `form:"type"`
	Content  string `form:"content"`
	Amount   string `form:"amount"`
	Symbol   string `form:"symbol"`
	QrValue  string `form:"qr_value" binding:"required"`
}
