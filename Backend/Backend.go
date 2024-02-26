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
		Action		 string `json:"action"`
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
	if(body.Action == "regist"){
		Services.Create(body.UserName, body.UserPassword)
	}

	// Create output when User is logged in
	var log string = ""
	if(body.Action == "login"){
		log, _ = Services.LoginCheck(body.UserName, body.UserPassword)
	}

	// Read all
	var users []Models.User= Services.Read()

	// ------------------------------------------- CONSOLE OUTPUT:	-------------------------------------------

	// Regist output
	if body.Action == "regist" {
		fmt.Println("--- REGIST ---")
		fmt.Println("Regist as Name:", body.UserName, "| Password:", body.UserPassword)

	}

	// Login output
	if body.Action == "login" {
		fmt.Println("--- LOGIN ---")
		fmt.Println("Login as Name:", body.UserName, "| Password:", body.UserPassword)
		fmt.Printf("%s\n", log)
	}

	fmt.Println()

	// Slice Output
	fmt.Println("Slice:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Hash: %s\n", user.Id, user.Name, user.Hash)
	}

	fmt.Println()

	// ------------------------------------------- Server output -------------------------------------------
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
