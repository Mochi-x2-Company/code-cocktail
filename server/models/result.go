package models

import (
	"github.com/jinzhu/gorm"
)

type Result struct {
	gorm.Model
	QuestionId int    `json: queston_id`
	UserID     int    `json: user_id`
	Code       string `json: code`
	Language   string `json: language`
}

func ResSingle() Result {
	return Result{}
}

func ResAll() []Result {
	return []Result{}
}

func ResMigrate() *Result {
	return &Result{}
}
