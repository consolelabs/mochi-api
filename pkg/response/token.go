package response

type Token struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Decimal int    `json:"decimal"`
}

type FindTokenByContractAddressResponse struct {
	Data Token `json:"data"`
}
