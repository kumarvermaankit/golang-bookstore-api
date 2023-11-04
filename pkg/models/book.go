package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kumarvermaankit/goBookstore/pkg/config"
	// "github.com/kumarvermaankit/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b

}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var Book []Book

	db := db.Where("id=?", id).Find(&Book)

	if len(Book) > 0 {
		return &Book[0], db

	}

	return nil, db
}

func DeleteBook(id int64) *Book {
	var book *Book
	db.Where("id=?", id).Delete(book)

	return book
}
