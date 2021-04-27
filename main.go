package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"

	"fmt"
	"github.com/crafter76/newmod"
	"example.com/m"
)

//var items []Item

type myHandler struct {

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

}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Структура объединяющая handler-функции`))
}

func main() {
	fmt.Println(newmod.Hi("Василий"))
	r := mux.NewRouter()
	items = append(items, Item{ID: "1", Name: "Георгин", Description: "Цветок"})
	items = append(items, Item{ID: "2", Name: "Пион", Description: "Цветок"})
	http.Handle("/", myHandler{})
	//r.HandleFunc("/items", getItem).Methods("GET")
	//r.HandleFunc("/items/{id}", getItem).Methods("GET")
	//r.HandleFunc("/items", postItem).Methods("POST")
	//r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))