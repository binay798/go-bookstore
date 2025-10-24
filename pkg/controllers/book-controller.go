package controllers

import (
	"encoding/json"
	"net/http"

	// "fmt"
	// "github.com/gorilla/mux"
	// "strconv"
	// "github.com/binay798/go-bookstore/pkg/utils"
	"github.com/binay798/go-bookstore/pkg/models"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {}
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBookById(w http.ResponseWriter, r *http.Request) {}
func UpdateBook(w http.ResponseWriter, r *http.Request)  {}
func DeleteBook(w http.ResponseWriter, r *http.Request)  {}
