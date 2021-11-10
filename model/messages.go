package model

import "time"

type messages struct {
	content     string
	create_time time.Time
	id          int64
	title       string
}
