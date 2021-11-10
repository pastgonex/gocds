package model

import "time"

type groups struct {
	create_time time.Time
	description string
	group_id    int64
	group_name  string
}
