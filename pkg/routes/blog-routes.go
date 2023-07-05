package routes

import (
	"github.com/adirn26/Blogpost-API/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBlogRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/blogs", controllers.GetBlog).Methods("GET")
	router.HandleFunc("/api/blogs/{id}", controllers.GetBlogByID).Methods("GET")
	router.HandleFunc("/api/blogs", controllers.CreateBlog).Methods("POST")
	router.HandleFunc("/api/blogs/{id}", controllers.UpdateBlog).Methods("PUT")
	router.HandleFunc("/api/blogs/{id}", controllers.DeleteBlog).Methods("DELETE")
}
