package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Contest struct {
	gorm.Model
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
