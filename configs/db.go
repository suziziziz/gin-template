package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB(dsn string) *gorm.DB {
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})

	db.AutoMigrate()

	return db
}
