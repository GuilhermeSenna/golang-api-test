package models

import (
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	Id int `json:"ID" gorm:"primary_key"`
	First string `json:"first"`
	Last string `json:"last"`
	Email string `json:"email"`
	Password string `json:"password"`
}