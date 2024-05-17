package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID     string `json:"id"`
	CMND   string `json:"cmnd"`
	Name   string `json:"name"`
	DiaChi string `json:"diaChi"`
}

var roll []Roll

// index
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roll)
}

// show
func getRoll(w http.ResponseWriter, r *http.Request) {

}

// crate
func createRoll(w http.ResponseWriter, r *http.Request) {

}

// delete
func updateRoll(w http.ResponseWriter, r *http.Request) {

}

// update
func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// khoi tao route
	route := mux.NewRouter()

	// endpoint
	router.HandleFunc("/tainn", getRolls).Methods("GET")
	router.HandleFunc("/tainn/{id}", getRoll).Methods("GET")
	router.HandleFunc("/tainn", createRoll).Methods("POST")
	router.HandleFunc("/tainn/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/tainn/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
