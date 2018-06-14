package main

import (
	"github.com/gorilla/mux"
	"github.com/vladmashtak/restapi/controllers"
	"log"
	"net/http"
)

func main() {
	const port string = ":8888"
	var bookController *controllers.BookController = controllers.GetBookController()
	var router *mux.Router = mux.NewRouter()

	router.HandleFunc("/api/books", bookController.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", bookController.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books", bookController.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id}", bookController.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books", bookController.DeleteBook).Methods(http.MethodDelete)

	log.Println("Server start at port " + port)
	log.Fatal(http.ListenAndServe(port, router))
}