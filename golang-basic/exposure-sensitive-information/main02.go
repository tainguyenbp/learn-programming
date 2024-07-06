package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Access sensitive data from the database
	// username := "admin"
	// password := "secret"

	// Instead of returning sensitive information, return a generic message
	fmt.Fprint(w, "Access denied")
	log.Printf("Received request for /users from Access denied", r.RemoteAddr)

}
