package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);not null"`
	FullName string `gorm:"type:varchar(100);not null"`
}

func InitDB(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
