package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	Input      string
	Output     string
	QuestionId int `gorm: "not null"`
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
