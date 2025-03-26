package main

import (
	"fmt"
	"my-book-api/handlers"
	"my-book-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
        if err := utils.LoadBooks(); err != nil {
                fmt.Println("Error loading books:", err)
                return
        }
         fmt.Println(utils.Books)
        r := mux.NewRouter()

        r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
        r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
		r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
                
		r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
		r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
		r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")
       

        fmt.Println("Starting server on port 8000...")
        if err := http.ListenAndServe(":8000", r); err != nil {
                fmt.Println("Error starting server:", err)
        }
}
