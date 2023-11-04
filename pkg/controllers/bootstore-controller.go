package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kumarvermaankit/goBookstore/pkg/models"
	"github.com/kumarvermaankit/goBookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	NewBook := &models.Book{}

	utils.ParseBody(r, NewBook)
	w.Header().Set("Content-Type", "pkglication/json")
	res := models.CreateBook(NewBook)

	val, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(val)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	Books := models.GetBooks()

	val, _ := json.Marshal(Books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(val)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		log.Fatal("Error occured")
	}

	b, _ := models.GetBook(ID)
	fmt.Println(ID)
	fmt.Println(b)
	val, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(val)

}

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Hello")

// }

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		log.Fatal("Error occured")
	}
	b := models.DeleteBook(ID)
	val, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(val)

}
