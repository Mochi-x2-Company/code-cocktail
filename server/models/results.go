package models

import ()

type Result struct {
	Id         int    `json:id sql:AUTO_INCREMENT`
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
