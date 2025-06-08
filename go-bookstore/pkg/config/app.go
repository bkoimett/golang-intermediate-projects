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
	d, err := gorm.Open("mysql", "bkoimett:bwanaChairman@tcp(127.0.0.1:3307)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

// used everywhere else
func GetDB() *gorm.DB{
	return db
}