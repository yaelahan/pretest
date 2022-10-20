package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"pretest-indihomesmart/models"
)

var db *gorm.DB

func NewDB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbUri := os.Getenv("DATABASE_URI")
	db, err = gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate db")
	}

	return db
}
