package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Migrate(db *gorm.DB) {
	que := QueMigrate()
	ans := AnsMigrate()
	cate := CateMigrate()
	res := ResMigrate()
	db.AutoMigrate(que, ans, cate, res)
	fmt.Println("Auto Migrate")
}
