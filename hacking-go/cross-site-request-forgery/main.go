package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	templates = template.Must(template.ParseFiles("index.html"))
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/transfer", transferHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == http.MethodPost {
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
