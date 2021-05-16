package main

import (
	"context"
	t "log"
	"time"

	"./controllers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mc := controllers.NewMqttController(getSession())
	go mc.GetSensorVal()
}
func getSession() *mongo.Session {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	session, err := client.StartSession()
	if err != nil {
		t.Fatal(err)
	}
	return &session
}
