package response

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseSucess struct {
	Success bool `json:"success"`
}

type ResponseStatus struct {
	Status string `json:"status"`
}
