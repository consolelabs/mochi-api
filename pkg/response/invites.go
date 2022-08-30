package response

import "github.com/defipod/mochi/pkg/model"

type GetInvitesResponse struct {
	Data []string `json:"data"`
}

type GetInviteTrackerConfigResponse struct {
	Data    *model.GuildConfigInviteTracker `json:"data"`
	Message string                          `json:"message"`
}

type ConfigureInvitesResponse struct {
	Data string `json:"data"`
}

type InvitesAggregationResponse struct {
	Data *UserInvitesAggregation `json:"data"`
}
