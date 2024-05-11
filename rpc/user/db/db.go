package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

func Init(connString string) {

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	//db = db.Debug()

	if err != nil {
		fmt.Println("storage err: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	autoMigrate(db)
	DB = db
}

func autoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(
		&Hello{},
	)
}
