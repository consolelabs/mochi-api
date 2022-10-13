package questuserreward

import "time"

type ListQuery struct {
	UserID    *string
	StartTime *time.Time
}
