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
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Payment{})

	initializeData(db)
}

func DbManager() *gorm.DB {
	return db
}

func initializeData(db *gorm.DB) {

	user_1 := model.User{
		Username: "Robert",
		Email: "robert@domain.com",
	}
	db.Create(&user_1)

	user_2 := model.User{
		Username: "Tomasz",
		Email: "tomasz@domain.com",
	}
	db.Create(&user_2)

	product_1 := model.Product{
		CategoryId: 1,
		Name: "iPhone 13 Pro",
		Price: 999.0,
	}
	db.Create(&product_1)

	product_2 := model.Product{
		CategoryId: 1,
		Name: "iPhone 12 Mini",
		Price: 599.0,
	}
	db.Create(&product_2)

	product_3 := model.Product{
		CategoryId: 1,
		Name: "iPhone SE",
		Price: 429.0,
	}
	db.Create(&product_3)

	category_1 := model.Category{
		Name: "Smartphone",
		Products: []model.Product{product_1, product_2, product_3},
	}
	db.Create(&category_1)

	payment_1 := model.Payment{
		OrderId: 1,
		Accepted: true,
		PaymentMethod: 1,
	}
	db.Create(&payment_1)

	order := model.Order{
		UserId: 1,
		Payment: payment_1,
		Products: []model.Product{product_1, product_2},
	}
	db.Create(&order)
}