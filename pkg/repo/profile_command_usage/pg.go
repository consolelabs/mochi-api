package profilecommandusage

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetTopProfileUsage(top int) ([]model.CommandUsageCounter, error) {
	q1 := pg.db.Table("profile_command_usages").Select("user_platform_id, profile_id, count(*) as usage").Where("command = '/profile' AND params = ''").Group("user_platform_id, profile_id")
	q2 := pg.db.Table("profile_command_usages").Select("replace(params, 'user:', '') as user_platform_id, count(*) as usage").Where("command = '/profile' AND params != ''").Group("params")

	rows, err := pg.db.Raw("WITH t1 AS (?), t2 AS (?) SELECT t1.usage + coalesce(t2.usage, 0) as total_usage, t1.user_platform_id, t1.profile_id FROM t1 left join t2 on t1.user_platform_id = t2.user_platform_id ORDER BY total_usage DESC LIMIT ?", q1, q2, top).Rows()
	if err != nil {
		return nil, err
	}

	var list []model.CommandUsageCounter
	for rows.Next() {
		var r model.CommandUsageCounter
		if err := rows.Scan(&r.TotalUsage, &r.UserPlatformId, &r.ProfileId); err != nil {
			return nil, err
		}
		list = append(list, r)
	}

	return list, nil
}
