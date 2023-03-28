package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bontl3/go-bookstore/pkg/models"
	"github.com/Bontl3/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Failed to marshal books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookId(w http.ResponseWriter, r *http.Request) {
	// In a URL, query parameters are used to pass data to a server as key-value pairs.
	// r.URL.Query().Get() returns value of a specific query parameter
	bookId := r.URL.Query().Get("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	bookDetails, _, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Failed to get book details", http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal book details", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookId(w http.ResponseWriter, r *http.Request) {
	bookId := r.URL.Query().Get("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		http.Error(w, "Failed to get book details", http.StatusBadRequest)
		return
	}

	deletedBookDetails, err := models.DeleteBook(ID)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(deletedBookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal book details", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	bookId := r.URL.Query().Get("bookId")

	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		http.Error(w, "Failed to update book", http.StatusBadRequest)
		return
	}

	bookDetails, db, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Can not get book details", http.StatusBadRequest)
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal book details", http.StatusBadRequest)
		return
	}

	db.Save(&bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
