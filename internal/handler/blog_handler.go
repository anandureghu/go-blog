package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/anandureghu/go-blog/internal/middleware"
	"github.com/anandureghu/go-blog/internal/model"
	"github.com/anandureghu/go-blog/internal/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	br repository.BlogRepository = *repository.NewBlogRepository()
)

func GetAllBlogs(w http.ResponseWriter, req *http.Request) {
	blogs, err := br.GetAllBlogs()
	if err != nil {
		http.Error(w, "internal server error", 500)
	}
	err = json.NewEncoder(w).Encode(blogs)
	w.WriteHeader(200)
}

func GetBlog(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid id", 400)
	}
	// Getting blog from repository
	blog, err := br.GetBlog(id)

	if err != nil {
		http.Error(w, "internal server error", 500)
	}
	if blog.Id == 0 {
		http.Error(w, fmt.Sprintf("blog with id %v not found", id), 404)
		return
	}

	json.NewEncoder(w).Encode(blog)
	w.WriteHeader(200)
}

func CreateBlog(w http.ResponseWriter, req *http.Request) {
	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		log.Println("Cannot unmarshal json")
		http.Error(w, "internal server error", 500)
	}

	err = middleware.ValidateBlog(blog)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = br.CreateBlog(blog)
	if err != nil {
		http.Error(w, "cannot create blog, internal server error", 500)
	}
	w.WriteHeader(201)
}

func UpdateBlog(w http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		log.Println("Cannot unmarshal json")
		http.Error(w, "internal server error", 500)
	}

	// updating blog from repository
	ub, err := br.UpdateBlog(id, blog)
	// Not found condition
	if err != nil && ub.Id <= 0 {
		http.Error(w, err.Error(), 404)
	}

	if err != nil {
		http.Error(w, "internal server error", 500)
	}

	json.NewEncoder(w).Encode(ub)
	w.WriteHeader(200)
}

func DeleteBlog(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid id", 400)
	}

	err = br.DeleteBlog(id)
	if err != nil && strings.Contains(err.Error(), strconv.Itoa(id)) {
		http.Error(w, err.Error(), 404)
	}

	if err != nil {
		http.Error(w, "internal server error", 500)
	}

	w.WriteHeader(200)
}
