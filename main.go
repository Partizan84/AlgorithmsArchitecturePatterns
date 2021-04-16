package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID          (string) `json:"id"`
	Name        (string) `json:"name"`
	Description (string) `json:"description"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

func postItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(rand.Intn(1000000))
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
	r := mux.NewRouter()
	items = append(items, Item{ID: "1", Name: "Георгин", Description: "Цветок"})
	items = append(items, Item{ID: "2", Name: "Пион", Description: "Цветок"})
	r.HandleFunc("/items", getItem).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", postItem).Methods("POST")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
