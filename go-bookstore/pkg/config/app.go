package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// enables interactiom with db
var (
	db * gorm.DB
)

// opens connection to db
func Connect(){
	d, err := gorm.Open("mysql", "bkoimett:08062016Vanna#/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

// used everywhere else
func GetDB() *gorm.DB{
	return db
}