package model

import (
	"github.com/jinzhu/gorm"
)

//User table for DB migration
type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
}