// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

type Book struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Author    string `json:"author"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Book Store!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllBooks")
    json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, Book := range Books {
        if Book.Id == key {
            json.NewEncoder(w).Encode(Book)
        }
    }
}


func createNewBook(w http.ResponseWriter, r *http.Request) {  
    reqBody, _ := ioutil.ReadAll(r.Body)
    var book Book 
    json.Unmarshal(reqBody, &book)
    Books = append(Books, book)

    json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, book := range Books {
        if book.Id == id {
            Books = append(Books[:index], Books[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/books", returnAllBooks)
    myRouter.HandleFunc("/book", createNewBook).Methods("POST")
    myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
    myRouter.HandleFunc("/book/{id}", returnSingleBook)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Books = []Book{
        Book{Id: "1", Title: "Hello", Author: "Book Author"},
        Book{Id: "2", Title: "Hello 2", Author: "Book Author"},
    }
    handleRequests()
}