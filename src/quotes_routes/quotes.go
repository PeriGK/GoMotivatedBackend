package quotes_routes

import (
	"github.com/gorilla/mux"
	"quotes/src/quotes_controllers"
)

func QuoteRequests(router *mux.Router) {
	router.HandleFunc("/quote/next", quotes_controllers.NextQuote).Methods("GET")
}
