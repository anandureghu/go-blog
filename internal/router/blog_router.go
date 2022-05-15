package router

import (
	"github.com/anandureghu/go-blog/internal/handler"
	"github.com/gorilla/mux"
)

func RouteBlog(r *mux.Router) {
	r.HandleFunc("/blogs", handler.GetAllBlogs).Methods("GET", "OPTIONS")
	r.HandleFunc("/blogs/{id}", handler.GetBlog).Methods("GET", "OPTIONS")
	r.HandleFunc("/blogs", handler.CreateBlog).Methods("POST", "OPTIONS")
	r.HandleFunc("/blogs/{id}", handler.UpdateBlog).Methods("PUT", "OPTIONS")
	r.HandleFunc("/blogs/{id}", handler.DeleteBlog).Methods("DELETE", "OPTIONS")
}
