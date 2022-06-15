package database

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {

	//Open sqlite connection and store data in ebiznes_go.db file
	db, err = gorm.Open("sqlite3", "ebiznes_go.db")

	if err != nil {
		panic("Error while connecting to DB")
	}

	//Migrate User table
	db.AutoMigrate(&model.User{})
}

func DbManager() *gorm.DB {
	return db
}