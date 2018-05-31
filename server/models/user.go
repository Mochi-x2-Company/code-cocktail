package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Provider   string
	Uid        string
	UserName   string
	ScreenName string
	ImageUrl   string
	IsAdmin    *bool `gorm: "default: false"`
}

func UserSingle() User {
	return User{}
}

func UserAll() []User {
	return []User{}
}

func UserMigrate() *User {
	return &User{}
}
