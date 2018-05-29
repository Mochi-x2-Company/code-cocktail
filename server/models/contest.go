package models

import (
	"time"
)

type Contest struct {
	Id        int       `json:id sql:AUTO_INCREMENT`
	Title     string    `json: title`
	StartTime time.Time `json: start_time`
	EndTime   time.Time `json: end_time`
}

func ConSingle() Contest {
	return Contest{}
}

func ConAll() []Contest {
	return []Contest{}
}

func ConMigrate() *Contest {
	return &Contest{}
}
