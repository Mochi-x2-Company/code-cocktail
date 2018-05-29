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

	db.AutoMigrate(que, ans, cate, res, team)
	fmt.Println("Auto Migrate")
}
