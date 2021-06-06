package main

import (
	"context"

	"net/http"

	"time"

	"Rest_api/controllers"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	router := mux.NewRouter()
	uc := controllers.NewUserController(client)
	router.HandleFunc("/case/{id}", uc.GetCase).Methods("GET","OPTIONS") 
	go router.HandleFunc("/case", uc.GetAllCases).Methods("GET", "OPTIONS")

	go controllers.GetSensorVal(client)

	http.ListenAndServe("localhost:8083", router)
}
