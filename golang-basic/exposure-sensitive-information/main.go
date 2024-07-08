// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/users", getUsers)
// 	http.ListenAndServe(":8080", nil)
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	// Access sensitive data from the database
// 	username := "admin"
// 	password := "secret"

// 	// Return the sensitive information in the HTTP response
// 	fmt.Fprintf(w, "Username: %s, Password: %s", username, password)
// }

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
	username := "admin"
	password := "secret"

	// Log the request
	log.Printf("Received request for /users from %s", r.RemoteAddr)

	// Return the sensitive information in the HTTP response
	fmt.Fprintf(w, "Username: %s, Password: %s", username, password)
}
