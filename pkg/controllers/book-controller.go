package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "fmt"
	"strconv"

	"github.com/gorilla/mux"

	// "github.com/binay798/go-bookstore/pkg/utils"
	"github.com/binay798/go-bookstore/pkg/models"
	"github.com/binay798/go-bookstore/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.Write([]byte("Failure"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	x, _ := models.GetBookById(bookId)
	details, _ := json.Marshal(x)
	fmt.Println(details)

	w.Write(details)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, _ := strconv.ParseInt(vars["bookId"], 0, 0)
	UpdatedBook := &models.Book{}
	utils.ParseBody(r, UpdatedBook)
	models.UpdateBook(bookId, *UpdatedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully updated book"))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		log.Fatal("Invalid book id")
		return
	}
	fmt.Println(bookId)
	models.DeleteBook(bookId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted book"))

}
