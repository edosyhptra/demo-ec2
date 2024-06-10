package main

import (
	"encoding/json"
	"net/http"
)

// Response struct to hold JSON response
type Response struct {
	Message string `json:"message"`
}

type Data struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{Name: "Sample Data", Value: 42}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/data", dataHandler)

	http.ListenAndServe(":8080", nil)
}
