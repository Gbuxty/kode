package main

import (
    "log"
	"net/http"

	"github.com/gorilla/mux"
	"kode/handlers"

)
   

	
func main(){
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/notes", handlers.GetNotesHandler).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNoteHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}