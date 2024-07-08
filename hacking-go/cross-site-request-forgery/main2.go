package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

var (
	templates = template.Must(template.ParseFiles("index.html"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/transfer", transferHandler)
	log.Fatal(http.ListenAndServe(":8080", csrf.Protect([]byte("32-byte-long-auth-key"))(nil)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		token := csrf.Token(r)
		data := struct {
			Token string
		}{
			Token: token,
		}
		templates.ExecuteTemplate(w, "index.html", data)
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Validate CSRF token
		if err := csrf.Protect([]byte("32-byte-long-auth-key")).VerifyToken(csrf.Token(r)); err != nil {
			http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
		}

		amount := r.FormValue("amount")
		account := r.FormValue("account")

		// Perform the money transfer
		if transferMoney(amount, account) {
			fmt.Fprintln(w, "Transfer successful!")
		} else {
			fmt.Fprintln(w, "Transfer failed!")
		}
	}
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	// Process transfer request
	// ...
}

func transferMoney(amount, account string) bool {
	// Perform money transfer logic
	// ...
	return false
}
