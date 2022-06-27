package model

import (
	"github.com/jinzhu/gorm"
)

//Product table for DB migration
type Product struct {
	gorm.Model
	ProductId int
	Name      string
	Price     float32
}