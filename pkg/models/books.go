package models

import (
	"example.com/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	NAME        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("id = ?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(book)
	return book
}

// update we can use create get and delete from controller
