package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/j-hanko/BookManagmentSystem-GO/pkg/models"
	"github.com/j-hanko/BookManagmentSystem-GO/pkg/utils"
	"net/http"
	"strconv"
)

var newBook models.Book

func getBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.getAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request) {

}
