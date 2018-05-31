package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	UserId    int `json: user_id`
	ContestId int `json: contest_id`
}

func TeamSingle() Team {
	return Team{}
}

func TeamAll() []Team {
	return []Team{}
}

func TeamMigrate() *Team {
	return &Team{}
}
