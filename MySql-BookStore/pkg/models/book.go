package models

import (
	"github.com/jinzhu/gorm"
	"github.com/yash-raj10/GO-Lang-Prjs/MySql-BookStore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
