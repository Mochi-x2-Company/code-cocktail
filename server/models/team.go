package models

import ()

type Team struct {
	Id        int `json:id sql:AUTO_INCREMENT`
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
