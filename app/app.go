package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//Declare endpoints
	router := mux.NewRouter()
	router.HandleFunc("/greeting", greet).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
