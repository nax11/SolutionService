package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nax11/solution_service/handler"
)

func main() {
	buildRoute()
}

func buildRoute() {
	router := mux.NewRouter()
	router.HandleFunc("/task/{taskName}", handler.Perform).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8090", router))
}
