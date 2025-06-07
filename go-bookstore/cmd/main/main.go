package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/bkoimett/golang-intermediate-projects/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes. 
}