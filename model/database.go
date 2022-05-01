package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Println(err)
	}

	return db, err

}
