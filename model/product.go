package model

import (
	"github.com/jinzhu/gorm"
)

//Product table for DB migration
type Product struct {
	gorm.Model
	Name      string
	Price     float32
}