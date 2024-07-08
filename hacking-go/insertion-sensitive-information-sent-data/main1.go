// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/login", login)
// 	log.Println("Starting server on :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	// Authenticate the user
// 	if !authenticate(username, password) {
// 		errMsg := fmt.Sprintf("Login failed for user: %s", username)
// 		log.Println(errMsg)
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	// Proceed with successful login
// 	successMsg := fmt.Sprintf("Welcome, %s!", username)
// 	tmpl, err := template.New("success").Parse("<h1>{{.}}</h1>")
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, successMsg)
// }

// func authenticate(username, password string) bool {
// 	// Perform authentication logic
// 	// Replace this with real authentication code
// 	if username == "admin" && password == "password" {
// 		return true
// 	}
// 	return false
// }


package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/login", login)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate the user
	if !authenticate(username, password) {
		errMsg := fmt.Sprintf("Login failed for user: %s", username)
		log.Println(errMsg)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Proceed with successful login
	successMsg := fmt.Sprintf("Welcome, %s!", username)
	tmpl, err := template.New("success").Parse("<h1>{{.}}</h1>")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, successMsg)
}

func authenticate(username, password string) bool {
	// Perform authentication logic
	// Replace this with real authentication code
	if username == "admin" && password == "password" {
		return true
	}
	return false
}
