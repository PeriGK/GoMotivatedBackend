package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"quotes/src/quotes_routes"
)

func main() {
	fmt.Print("Starting GoMotivated")

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	quotes_routes.QuoteRequests(router)
	handler := handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}