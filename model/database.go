package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func Database() (*gorm.DB, error) {

	DB, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	DB.AutoMigrate(&Grocery{})

	return DB, err

}
