package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	_db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := _db.AutoMigrate(&Grocery{}); err != nil {
		panic(err)
	}

	Db := _db
	return Db
}
