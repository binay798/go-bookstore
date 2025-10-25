package models

import (
	"github.com/binay798/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json: "name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

func UpdateBook(ID int64, updatedBook Book) Book {
	db.Model(&Book{}).Where("id = ?", ID).Updates(updatedBook)

	// db.Save(book)
	// book, _ := GetBookById(ID)
	// res, _ := json.Marshal(book)
	// fmt.Println(string(res))

	// createdAt := book.CreatedAt
	// updatedAt := book.UpdatedAt
	// book = updatedBook
	// book.CreatedAt = createdAt
	// book.UpdatedAt = updatedAt
	// fmt.Println(book)
	// db.Save(book)

	return updatedBook
}
