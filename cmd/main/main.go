package main

import (
	"log"
	"net/http"

	"github.com/adirn26/Blogpost-API/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBlogRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
