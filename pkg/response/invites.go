package response

type GetInvitesResponse struct {
	Data []string `json:"data"`
}

type ConfigureInvitesResponse struct {
	Data string `json:"data"`
}

type InvitesAggregationResponse struct {
	Data *UserInvitesAggregation `json:"data"`
}
