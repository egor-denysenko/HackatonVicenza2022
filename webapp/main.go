package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, map[string]any{"prova": "testoprova"})
}

func handleRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles("public/stanza.html")
	t.Execute(w, map[string]any{"id": vars["roomId"]})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleMain)
	r.HandleFunc("/stanze/{roomId}", handleRoom)

	log.Fatal(http.ListenAndServe(":8000", r))
}
