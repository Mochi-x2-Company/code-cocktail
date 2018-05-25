package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Migrate(db *gorm.DB) {
	que := QueMigrate()
	db.AutoMigrate(que)
	fmt.Println("Auto Migrate")
}
