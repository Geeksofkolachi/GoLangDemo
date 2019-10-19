package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,pmitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,pmitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	collection := client.Database("Application").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)

}

func GetUserEndpoint(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println(id)
	var user User
	collection := client.Database("Application").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var users []User
	collection := client.Database("Application").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person User
		cursor.Decode(&person)
		users = append(users, person)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")

	http.ListenAndServe(":8000", router)
}
