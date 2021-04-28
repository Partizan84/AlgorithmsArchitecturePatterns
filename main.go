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
)

//var items []Item



func main() {
	fmt.Println(newmod.Hi("Василий"))
	r := mux.NewRouter()
	items = append(items, Item{ID: "1", Name: "Георгин", Description: "Цветок"})
	items = append(items, Item{ID: "2", Name: "Пион", Description: "Цветок"})
	http.Handle("/", myHandler{})
	log.Fatal(http.ListenAndServe(":8080", r))
}

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))