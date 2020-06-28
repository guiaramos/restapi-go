package  main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book Struct (Model)

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Inital Books
var books []Book

// Controllers

func getBooks(res http.ResponseWriter, req *http.Request){

}

func getBook(res http.ResponseWriter, req *http.Request){

}

func createBook(res http.ResponseWriter, req *http.Request){

}

func updateBook(res http.ResponseWriter, req *http.Request){

}

func deleteBook(res http.ResponseWriter, req *http.Request){

}


func main(){
	// Assign Router
	router := mux.NewRouter()

	// Mock data - @TODO: Implement DB
	books = append(books, Book{ID: "1", Isbn: "23423", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "45453", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})


	// Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Listen
	log.Fatal(http.ListenAndServe(":8000",router))
}


