package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"response": "pong"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/ping", PingHandler).Methods("GET")

    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
