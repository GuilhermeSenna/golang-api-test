package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	First string `json:"first"`
	Last string `json:"last"`
	Email string `json:"email"`
	Password string `json:"password"`
}