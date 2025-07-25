package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.Handle("/terminal", ValidateTokenMiddleware(http.HandlerFunc(CreateTerminalHandler))).Methods("POST")

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
