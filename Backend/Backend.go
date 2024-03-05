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

	// Close body at the end
	defer r.Body.Close()

	if r.Method == "POST" {
		// POST request handling for registration or login
		var body struct {
			Action       string `json:"action"`
			UserName     string `json:"userName"`
			UserPassword string `json:"userPassword"`
		}

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "ERROR (golang): ", http.StatusBadRequest)
			return
		}
		
		fmt.Printf("Received %s request from %s with body: %v\n", r.Method, r.RemoteAddr, body)

		// Create output when user is regist
		var checkReg bool

		if body.Action == "regist" {
			checkReg = Services.RegCheck(body.UserName)
			if checkReg {
				Services.Create(body.UserName, body.UserPassword)
			}
		}

		// Create output when user is logged in
		var checkLog bool

		if body.Action == "login" {
			checkLog = Services.LoginCheck(body.UserName, body.UserPassword)
		}

		// Read all
		var users []Models.User = Services.Read()

		// ------------------------------------------- CONSOLE OUTPUT:	-------------------------------------------

		// Regist output
		if body.Action == "regist" {
			fmt.Println("--- REGIST ---")
			if checkReg {
				fmt.Println("Regist: OK")
				fmt.Println("Regist as Name:", body.UserName, "| Password:", body.UserPassword)
			} else {
				fmt.Println("Regist: FALSE")
				fmt.Println("Name:", body.UserName, "| Password:", body.UserPassword)
			}
		}

		// Login output
		if body.Action == "login" {
			fmt.Println("--- LOGIN ---")
			if checkLog {
				fmt.Println("Login: OK")
				fmt.Println("Login as Name:", body.UserName, "| Password:", body.UserPassword)
			} else {
				fmt.Println("Login: FALSE")
				fmt.Println("Name:", body.UserName, "| Password:", body.UserPassword)
			}
		}

		fmt.Println()

		// Slice Output
		fmt.Println("Slice:")
		for _, user := range users {
			fmt.Printf("ID: %d, Name: %s, Hash: %s\n", user.Id, user.Name, user.Hash)
		}

		fmt.Println()

		// ------------------------------------------- Server output -------------------------------------------
		fmt.Fprintln(w, "Golang-Server:")
		// fmt.Fprintln(w, "> JSON:", body)

	} else if r.Method == "GET" {
		// GET request handling for other purposes
		// ...
	}
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
