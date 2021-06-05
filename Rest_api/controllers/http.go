package controllers

import (
	"Rest_api/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetController struct {
	mongoclient *mongo.Client
}

func NewUserController(c *mongo.Client) *GetController {
	return &GetController{c}
}

func (gc GetController) GetCase(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id := params["id"]
	fmt.Println(id)
	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cased := models.AlertMssg{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := gc.mongoclient.Database("fir").Collection("cases")
	if err := collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&cased); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if err := json.NewEncoder(w).Encode(cased); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func (gc GetController) GetAllCases(w http.ResponseWriter, req *http.Request) {

	// An array in which you can store the decoded documents
	cased := []models.AlertMssg{}

	// Pass these options to the Find method
	findOptions := options.Find()
	// findOptions.SetLimit(2)   //set the limit of documents to fetch

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := gc.mongoclient.Database("fir").Collection("cases")

	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded

		altmssg := models.AlertMssg{}
		err := cur.Decode(&altmssg)
		if err != nil {
			log.Fatal(err)
		}

		cased = append(cased, altmssg)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//convert to array of struct alertmssg into json the send through the wire as response
	if err := json.NewEncoder(w).Encode(cased); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
