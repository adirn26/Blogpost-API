package models

import (
	"github.com/adirn26/Blogpost-API/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	Body     string `json:"body"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Blog{})
}

func (b *Blog) CreateBlog() *Blog {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBlogs() []Blog {
	var Blogs []Blog
	db.Find(&Blogs)
	return Blogs
}

func GetBlogById(Id int64) (*Blog, *gorm.DB) {
	var getBlog Blog
	db := db.Where("ID = ?", Id).Find(&getBlog)
	return &getBlog, db
}

func DeleteBlogById(Id int64) Blog {
	var Blog Blog
	db.Where("ID = ?", Id).Delete(Blog)
	return Blog
}

func GetByCategory(category string) []Blog {
	var Blogs []Blog
	db.Where("category = ?", category).Find(&Blogs)
	return Blogs
}
