package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adirn26/Blogpost-API/pkg/models"
	"github.com/adirn26/Blogpost-API/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBlog models.Blog

// CreateBlog ... Create Blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	CreateBlog := &models.Blog{}
	utils.ParseBody(r, CreateBlog)
	blog := CreateBlog.CreateBlog()
	res, _ := json.Marshal(blog)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBlog ... Get all Blogs
func GetBlog(w http.ResponseWriter, r *http.Request) {
	Blogs := models.GetAllBlogs()
	res, _ := json.Marshal(Blogs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBlogByID ... Get the Blog by id
func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BlogID := vars["id"]
	ID, err := strconv.ParseInt(BlogID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	BlogDetails, _ := models.GetBlogById(ID)
	res, _ := json.Marshal(BlogDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBlog ... Update the Blog
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var updateBlog = &models.Blog{}
	utils.ParseBody(r, updateBlog)
	vars := mux.Vars(r)
	BlogID := vars["id"]
	ID, err := strconv.ParseInt(BlogID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	BlogDetails, db := models.GetBlogById(ID)
	if updateBlog.Title != "" {
		BlogDetails.Title = updateBlog.Title
	}
	if updateBlog.Body != "" {
		BlogDetails.Body = updateBlog.Body
	}
	if updateBlog.Author != "" {
		BlogDetails.Author = updateBlog.Author
	}
	if updateBlog.Category != "" {
		BlogDetails.Category = updateBlog.Category
	}
	db.Save(&BlogDetails)
	res, _ := json.Marshal(BlogDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBlog ... Delete the Blog
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BlogID := vars["id"]
	ID, err := strconv.ParseInt(BlogID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Blog := models.DeleteBlogById(ID)
	res, _ := json.Marshal(Blog)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Category := vars["category"]
	BlogDetails := models.GetByCategory(Category)
	res, _ := json.Marshal(BlogDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
