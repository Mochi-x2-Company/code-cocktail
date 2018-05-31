package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
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
