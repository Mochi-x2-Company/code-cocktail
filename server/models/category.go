package models

import (
	"time"
)

type Category struct {
	Id        int       `json:id sql:AUTO_INCREMENT`
	Title     string    `json: title`
	StartTime time.Time `json: start_time`
	EndTime   time.Time `json: end_time`
}

func CateSingle() Category {
	return Category{}
}

func CateAll() []Category {
	return []Category{}
}

func CateMigrate() *Category {
	return &Category{}
}
