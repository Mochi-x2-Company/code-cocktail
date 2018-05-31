package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	Input      string `json: input`
	Output     string `json: output`
	QuestionId int    `json: queston_id`
}

func AnsSingle() Answer {
	return Answer{}
}

func AnsAll() []Answer {
	return []Answer{}
}

func AnsMigrate() *Answer {
	return &Answer{}
}
