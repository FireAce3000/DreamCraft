package main

import (
	"DreamCraft/Models"
	"DreamCraft/Services"
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

	// JSON write in body object with error
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "ERROR (golang): "+err.Error(), http.StatusBadRequest)
		return
	}

	// Close body at the end
	defer r.Body.Close()

	// Add user
	Services.Create(body.UserName, body.UserPassword)

	// Read all
	var users []Models.User= Services.Read()

	// Console output
	fmt.Println("--- LOGIN ---")

	fmt.Println("New User: Name:", body.UserName, "| Password:", body.UserPassword)

	// Output
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Hash: %s\n", user.Id, user.Name, user.Hash)
	}

	// Server output
	// fmt.Fprintln(w, "Golang-Server:")
	// fmt.Fprintln(w, "> JSON:", body)
}

func main() {
	// Display the Server, when online
	fmt.Println("")
	fmt.Println("--- DreamCraft Server ---")
	fmt.Println("> url: http://localhost:8080/")
	fmt.Println("> v 0.0.1")
	fmt.Println("")

	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
