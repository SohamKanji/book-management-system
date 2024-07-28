package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SohamKanji/book-management-system/pkg/models"
	"github.com/SohamKanji/book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	new_books := models.GetAllBooks()
	res, _ := json.Marshal(new_books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id := vars["id"]
	id, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	book_details, _ := models.GetBookById(id)
	res, _ := json.Marshal(book_details)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	create_book := &models.Book{}
	utils.ParseBody(r, create_book)
	b := create_book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id := vars["id"]
	id, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	deleted_book := models.DeleteBook(id)
	res, _ := json.Marshal(deleted_book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id := vars["id"]
	id, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	update_book := &models.Book{}
	utils.ParseBody(r, update_book)
	book_details, db := models.GetBookById(id)
	if update_book.Name != "" {
		book_details.Name = update_book.Name
	}
	if update_book.Author != "" {
		book_details.Author = update_book.Author
	}
	if update_book.Publication != "" {
		book_details.Publication = update_book.Publication
	}
	db.Save(&book_details)
	res, _ := json.Marshal(book_details)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
