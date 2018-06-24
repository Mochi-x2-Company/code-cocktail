package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// 本番環境用データベースコネクション
func PQPro() *gorm.DB {
	DBMS := "postgres"
	USER := "code_cocktail"
	PROTOCOL := "localhost"
	DBNAME := "code_cocktail"
	PASS := "password" // 本番環境ではここの値を環境変数などに変更する

	db, err := gorm.Open(DBMS, "host="+PROTOCOL+" user="+USER+" dbname="+DBNAME+" password="+PASS)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// 開発環境用データベースコネクション
func PQDev() *gorm.DB {
	DBMS := "postgres"
	USER := "code_cocktail"
	PROTOCOL := "localhost"
	DBNAME := "code_cocktail_development"
	PASS := "password" // 本番環境ではここの値を環境変数などに変更する

	db, err := gorm.Open(DBMS, "host="+PROTOCOL+" user="+USER+" dbname="+DBNAME+" password="+PASS)

	if err != nil {
		panic(err.Error())
	}

	return db
}
