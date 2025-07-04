package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bkoimett/golang-intermediate-projects/go-bookstore/pkg/utils"
	"github.com/bkoimett/golang-intermediate-projects/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	response, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response) //helps us send sth to front end
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0) //book id is in string convert to int
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID) //gets us book hy ID
	res, _ := json.Marshal(bookDetails) //json version of book details
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)// status 200 - everything is fine
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)//change from json to a form that  db can understand
	b := CreateBook.CreateBook() //--magi db---
	res, _ := json.Marshal(b)//change to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	 ID, err := strconv.ParseInt(bookId, 0,0)
	 if err != nil {
		fmt.Println("error while parsing")
	 }

	 book := models.DeleteBook(ID)
	 res, _ := json.Marshal(book)
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)
}

// has aspects of all the other functions

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}

	//
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}