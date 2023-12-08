package routes

import (
	"github.com/gorilla/mux"
	"github.com/yash-raj10/GO-Lang-Prjs/MySql-BookStore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookId).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}