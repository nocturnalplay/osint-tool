package routers

import (
	"cybercops.in/controller"
	"github.com/gorilla/mux"
)

var Router = func(router *mux.Router) {
	router.HandleFunc("/", controller.Root).Methods("GET")
	router.HandleFunc("/ip-info",controller.Ipinfo).Methods("GET")
	router.HandleFunc("/domain-info",controller.DomainCheck).Methods("GET")
	router.HandleFunc("/search-key",controller.SearchKey).Methods("GET")
}
