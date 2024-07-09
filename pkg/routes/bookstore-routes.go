package routes

import (
	"example.com/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/bookstore/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/bookstore/books/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/bookstore/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/bookstore/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/bookstore/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
