package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web_template/config"
)

var Db *gorm.DB
var err error

func Connect() {
	Db, err = gorm.Open(mysql.Open(config.Config.DbString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := Db.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)

	db.SetMaxOpenConns(100)

	db.SetConnMaxLifetime(time.Hour)
}

func AutoMigrate() {
	Db.AutoMigrate()
}
