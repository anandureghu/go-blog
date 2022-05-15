package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anandureghu/go-blog/internal/middleware"
	"github.com/anandureghu/go-blog/internal/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	PORT, URL string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	URL = os.Getenv("URL")

}

func main() {

	r := mux.NewRouter()
	router.RouteBlog(r)
	r.Use(middleware.HeaderMiddleware)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	// headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	// originsOk := handlers.AllowedOrigins([]string{"http://localhost:*", "*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("Server started listening on %v%v/\n", URL, PORT)
	// log.Fatal(http.ListenAndServe(PORT, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
	log.Fatal(http.ListenAndServe(PORT, r))
	// defer repository.GetConnection().Close()
}
