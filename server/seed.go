package main

import (
	"./config"
	"./models"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

// Questionテーブルに値をinsertしていく
func que(db *gorm.DB) {
	for i := 1; i <= 100; i++ {
		que := models.QueSingle()
		que.Title = "hoge" + (strconv.Itoa(i))
		que.Body = "hogehogehoge" + (strconv.Itoa(i))
		db.Create(&que)
	}
}

// ダミーデータを挿入する
func seed(db *gorm.DB) {
	// 一応Migrateをしておく
	models.Migrate(db)

	// 各関数を呼び出してテーブルにinsertしていく
	que(db)

	fmt.Println("Insert seeds!!")
}

func main() {
	flag.Parse()

	// 実行時引数によって使用する環境を変える
	if flag.Arg(0) == "production" {
		db := config.PQPro()
		defer db.Close()
		seed(db)
	} else {
		db := config.PQDev()
		defer db.Close()
		seed(db)
	}
}