package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// Allow access from all sources
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow specific headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		// For preflight requests, allow access to specific methods
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		return
	}

	// Decode JSON body at the end
	var body struct {
		UserName     string `json:"userName"`
		UserPassword string `json:"userPassword"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "ERROR (golang): "+err.Error(), http.StatusBadRequest)
		return
	}

	// Close body at the end
	defer r.Body.Close()

	// Console output
	fmt.Println("--- LOGIN ---")
	fmt.Println("Name:", body.UserName, "| Password:", body.UserPassword)

	// Server output
	fmt.Fprintln(w, "Golang-Server:")
	fmt.Fprintln(w, "> JSON:", body)
}

func main() {
	fmt.Println("")
	fmt.Println("--- DreamCraft Server ---")
	fmt.Println("> url: http://localhost:8080/")
	fmt.Println("v 0.0.1")
	fmt.Println("")

	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
