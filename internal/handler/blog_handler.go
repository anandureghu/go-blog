package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/anandureghu/go-blog/internal/model"
	"github.com/anandureghu/go-blog/internal/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	br repository.BlogRepository = *repository.NewBlogRepository()
)

func GetAllBlogs(w http.ResponseWriter, req *http.Request) {
	blogs := br.GetAllBlogs()
	json.NewEncoder(w).Encode(blogs)
}

func GetBlog(w http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(405)
		fmt.Fprintf(w, "invalid id")
		return
	}
	blog, err := br.GetBlog(id)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "internal server error")
		return
	}
	if err == nil && blog.Id == 0 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "not found")
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func CreateBlog(w http.ResponseWriter, req *http.Request) {
	blog := model.Blog{}
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "internal server error")
		log.Println("Cannot unmarshal json")
	}
	w.WriteHeader(201)
	br.CreateBlog(blog)
}

func UpdateBlog(w http.ResponseWriter, req *http.Request) {

}

func DeleteBlog(w http.ResponseWriter, req *http.Request) {

	log.Println("---------------------DELETE----------------------")
	w.Header().Set("Access-Control-Request-Method", "DELETE")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// req.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	param := mux.Vars(req)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(405)
		fmt.Fprintf(w, "invalid id")
	}
	br.DeleteBlog(id)
}
