package model

import (
	"github.com/jinzhu/gorm"
)

//Order table for DB migration
type Order struct {
	gorm.Model
	UserId   int
	Payment  Payment
	Products []Product `gorm:"many2many:order_product;"`
}