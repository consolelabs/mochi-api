package invite_histories

import (
	"database/sql"
	"encoding/json"
	"sort"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(invite *model.InviteHistory) error {
	return pg.db.Create(invite).Error
}

func (pg *pg) CountByInviter(inviterID int64) (int64, error) {
	var count int64
	err := pg.db.Model(&model.InviteHistory{}).Where("inviter_id = ?", inviterID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (pg *pg) GetInvitesLeaderboard(guildID string) (resp []response.UserInvitesAggregation, err error) {
	rows, err := pg.db.Raw(`
		select
		aggr.invited_by, 
		jsonb_object_agg(
			aggr.type,
			aggr.amount
		) as invite_aggr,
		aggr.guild_id
		from(
			select 
				count("type") as amount,
				invited_by,
				"type",
				guild_id
			from
			(	
				with aggr_max as 
				(
					select 
						max(created_at) as created_at,
						user_id,
						invited_by,
						guild_id
					from invite_histories
					group by (
						user_id,
						invited_by,
						guild_id
					)
				)
				select 
					i.guild_id,
					i.user_id,
					i.invited_by,
					i.created_at,
					i.type
				from invite_histories as i
				inner join aggr_max as a
				on (
					i.created_at = a.created_at and
					i.guild_id = a.guild_id and
					i.user_id = a.user_id and 
					i.invited_by = a.invited_by
				)
			) as full_aggr_max
			group by
			(
				invited_by,
				"type",
				guild_id
			) 
		) as aggr
		where aggr.guild_id = ?
		group by (aggr.invited_by, aggr.guild_id)
	`, guildID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var invitedBy string
		var guildID string
		var inviteAggr []byte
		if err := rows.Scan(&invitedBy, &inviteAggr, &guildID); err != nil {
			return nil, err
		}

		var inviteAggrMap map[string]int
		if err := json.Unmarshal(inviteAggr, &inviteAggrMap); err != nil {
			return nil, err
		}

		regular := inviteAggrMap["normal"] + inviteAggrMap["fake"] + inviteAggrMap["left"]
		resp = append(resp, response.UserInvitesAggregation{
			InviterID: invitedBy,
			Regular:   regular,
			Fake:      inviteAggrMap["fake"],
			Left:      inviteAggrMap["left"],
		})
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Regular > resp[j].Regular
	})

	// Top 10
	if len(resp) > 10 {
		resp = resp[:10]
	}

	return resp, nil
}

func (pg *pg) GetUserInvitesAggregation(guildID, inviterID string) (*response.UserInvitesAggregation, error) {
	var invitedBy string
	var inviteGuildID string
	var inviteAggr []byte

	err := pg.db.Raw(`
	select
	aggr.invited_by, 
	jsonb_object_agg(
		aggr.type,
		aggr.amount
	) as invite_aggr,
	aggr.guild_id
	from(
		select 
			count("type") as amount,
			invited_by,
			"type",
			guild_id
		from
		(	
			with aggr_max as 
			(
				select 
					max(created_at) as created_at,
					user_id,
					invited_by,
					guild_id
				from invite_histories
				group by (
					user_id,
					invited_by,
					guild_id
				)
			)
			select 
				i.guild_id,
				i.user_id,
				i.invited_by,
				i.created_at,
				i.type
			from invite_histories as i
			inner join aggr_max as a
			on (
				i.created_at = a.created_at and
				i.guild_id = a.guild_id and
				i.user_id = a.user_id and 
				i.invited_by = a.invited_by
			)
		) as full_aggr_max
		group by
		(
			invited_by,
			"type",
			guild_id
		) 
	) as aggr
	where aggr.guild_id = ? and aggr.invited_by = ?
	group by (aggr.invited_by, aggr.guild_id)
`, guildID, inviterID).Row().Scan(&invitedBy, &inviteAggr, &inviteGuildID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return &response.UserInvitesAggregation{
			InviterID: inviterID,
		}, nil
	}

	var inviteAggrMap map[string]int
	if err := json.Unmarshal(inviteAggr, &inviteAggrMap); err != nil {
		return nil, err
	}
	regular := inviteAggrMap["normal"] + inviteAggrMap["fake"] + inviteAggrMap["left"]

	return &response.UserInvitesAggregation{
		InviterID: invitedBy,
		Regular:   regular,
		Fake:      inviteAggrMap["fake"],
		Left:      inviteAggrMap["left"],
	}, nil
}
