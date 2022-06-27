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

	//Migrate tables
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})

	initializeData(db)
}

func DbManager() *gorm.DB {
	return db
}

func initializeData(db *gorm.DB) {

	product_1 := model.Product{
		ProductId: 1,
		Name: "iPhone 13 Pro",
		Price: 999.0,
	}
	db.Create(&product_1)

	product_2 := model.Product{
		ProductId: 2,
		Name: "iPhone 12 Mini",
		Price: 599.0,
	}
	db.Create(&product_2)

	product_3 := model.Product{
		ProductId: 3,
		Name: "iPhone SE",
		Price: 429.0,
	}
	db.Create(&product_3)

	order := model.Order{
		OrderId: 1,
		Products: []model.Product{product_1, product_2},
	}
	db.Create(&order)
}