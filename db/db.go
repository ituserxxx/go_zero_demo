package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

func Init() {
	connString := "demo:demo@tcp(172.16.9.116:33061)/demo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
	db = db.Debug()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
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
