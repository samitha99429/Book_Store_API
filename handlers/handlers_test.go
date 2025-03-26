package handlers

import (

	"encoding/json"
	"my-book-api/models"
	"my-book-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
        utils.Books = []models.Book{
                {BookID: "1", Title: "Test Book 1"},
                {BookID: "2", Title: "Test Book 2"},
        }

        req, err := http.NewRequest("GET", "/books", nil)
        if err != nil {
                t.Fatal(err)
        }
        rr := httptest.NewRecorder()

        GetBooks(rr, req)

   
        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v",
                        status, http.StatusOK)
        }

        var got []models.Book
        err = json.Unmarshal(rr.Body.Bytes(), &got)
        if err != nil {
                t.Fatal(err)
        }

        if len(got) != 2 {
                t.Errorf("handler returned unexpected body: got %v want %v",
                        len(got), 2)
        }

        if got[0].Title != "Test Book 1" {
                t.Errorf("handler returned unexpected body: got %v want %v", got[0].Title, "Test Book 1")
        }
}