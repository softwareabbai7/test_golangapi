package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	message := fmt.Sprintf("%s, hello! How are you?", name)

	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"response": "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler).Methods("GET")
	r.HandleFunc("/greet/{name}", greetHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
