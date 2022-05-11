package router

import (
	"github.com/anandureghu/go-blog/internal/handler"
	"github.com/gorilla/mux"
)

func init() {

}

func RouteBlog(r *mux.Router) {
	r.HandleFunc("/blogs", handler.GetAllBlogs).Methods("GET")
	r.HandleFunc("/blogs/{id}", handler.GetBlog).Methods("GET")
	r.HandleFunc("/blogs", handler.CreateBlog).Methods("POST")
	r.HandleFunc("/blogs/{id}", handler.UpdateBlog).Methods("PUT")
	r.HandleFunc("/blogs/{id}", handler.DeleteBlog).Methods("DELETE", "OPTIONS")
}
