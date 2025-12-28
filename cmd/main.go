package main

import (
	"github.com/gorilla/mux"
	"github.com/j-hanko/BookManagmentSystem-GO/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
}
