
package handlers

import (
    "encoding/json"
    "fmt" 
    "my-book-api/models"
    "my-book-api/utils"
    "net/http"
    "strings"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-Type", "application/json")
    json.NewEncoder(w).Encode(utils.Books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book models.Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.BookID = uuid.New().String()
    utils.Books = append(utils.Books, book)
    _ = utils.SaveBooks()
    json.NewEncoder(w).Encode(book)
}





func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range utils.Books {
        if item.BookID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Book{})
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range utils.Books {
        if item.BookID == params["id"] {
            utils.Books = append(utils.Books[:index], utils.Books[index+1:]...)
            var book models.Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.BookID = params["id"]
            utils.Books = append(utils.Books, book)
            _ = utils.SaveBooks()
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(utils.Books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range utils.Books {
        if item.BookID == params["id"] {
            utils.Books = append(utils.Books[:index], utils.Books[index+1:]...)
            _ = utils.SaveBooks()
            break
        }
    }
    json.NewEncoder(w).Encode(utils.Books)
}



func SearchBooks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("SearchBooks handler called")
    w.Header().Set("Content-Type", "application/json")
    query := r.URL.Query().Get("q")
    fmt.Println("Search query:", query) 
    if query == "" {
        fmt.Println("Empty query, returning empty array") 
        json.NewEncoder(w).Encode([]models.Book{})
        return
    }

    results := searchBooksConcurrently(query)
    fmt.Println("Search results:", results) 
    json.NewEncoder(w).Encode(results)
}

func searchBooksConcurrently(query string) []models.Book {
    bookCount := len(utils.Books)
    if bookCount == 0 {
        return []models.Book{}
    }

    numGoroutines := 4
    chunkSize := (bookCount + numGoroutines - 1) / numGoroutines
    resultsChan := make(chan []models.Book, numGoroutines)

          for i := 0; i < numGoroutines; i++ {
          start := i * chunkSize
          end := start + chunkSize
        if end > bookCount {
            end = bookCount
        }

        go func(books []models.Book, q string, results chan<- []models.Book) {
            localResults := []models.Book{}
            for _, book := range books {
                fmt.Println("Checking book:", book.Title, book.Description) 
                if strings.Contains(strings.ToLower(book.Title), strings.ToLower(q)) ||
                    strings.Contains(strings.ToLower(book.Description), strings.ToLower(q)) {
                    localResults = append(localResults, book)
                }
            }
            fmt.Println("Goroutine results:", localResults) 
            results <- localResults
        }(utils.Books[start:end], query, resultsChan)
    }

      finalResults := []models.Book{}
    for i := 0; i < numGoroutines; i++ {
        finalResults = append(finalResults, <-resultsChan...)
    }
    fmt.Println("Final results:", finalResults) 
    return finalResults
}