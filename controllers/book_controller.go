package controllers

import (
	"encoding/json"
	"github.com/vladmashtak/restapi/datasource"
	"net/http"
)

type BookController struct {
}

var bookController *BookController = nil

func GetBookController() *BookController {
	if bookController == nil {
		bookController = &BookController{}
	}

	return bookController
}

// Get All Books
func (b *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datasource.Books)
}

// Get Single Book
func (b *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
}

// Create a New Book
func (b *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
}

// Update a Book
func (b *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
}

// Delete a Book
func (b *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
}
