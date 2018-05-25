package models

import ()

type Question struct {
	Id    int    `json:id sql:AUTO_INCREMENT`
	Title string `json: title`
	Body  string `json: body`
}

func QueSingle() Question {
	return Question{}
}

func QueAll() []Question {
	return []Question{}
}

func QueMigrate() *Question {
	return &Question{}
}
