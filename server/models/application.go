package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Migrate(db *gorm.DB) {
	que := QueMigrate()
	ans := AnsMigrate()
	cate := ConMigrate()
	res := ResMigrate()
	team := TeamMigrate()
	user := UserMigrate()

	db.AutoMigrate(que, ans, cate, res, team, user)
	fmt.Println("Auto Migrate")
}
