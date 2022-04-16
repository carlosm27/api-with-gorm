package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Database() {

	_db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := _db.AutoMigrate(&Grocery{}); err != nil {
		panic(err)
	}

	Db = _db

}
