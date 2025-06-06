package models

import (
	"github.com/jinzhu/gorm"
	"github.com/bkoimett/golang-intermediate-projects/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

