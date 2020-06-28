package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Inital Books
var books []Book

// Controllers

func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

func getBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(res).Encode(book)
			return
		}
	}

	json.NewEncoder(res).Encode(&Book{})

}

func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	json.NewEncoder(res).Encode(book)

}

func updateBook(res http.ResponseWriter, req *http.Request) {

}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, book := range books{
		if book.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(books)
}

func main() {
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
	log.Fatal(http.ListenAndServe(":8000", router))
}
