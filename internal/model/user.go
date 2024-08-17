package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Id        uint
	Username  string
	LastName  string
	FirstName string
	Email     string
	Password  string
	role      string
}
