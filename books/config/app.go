package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// Connect to the database
	d, err := gorm.Open("mysql", "CONNECTION_URL")
	if err != nil {
		log.Fatalln(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
