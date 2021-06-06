package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sensor struct {
	Accelerometer []float64 `json:"accelerometer" bson:"accelerometer"`
	Gyroscope     []float64 `json:"gyroscope" bson:"gyroscope"`
	Sound         int       `json:"sound" bson:"sound"`
}

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	PhNo     string             `json:"number"  bson:"number"`
	Loc      Location           `json:"location" bson:"location"`
	Endpoint string             `json:"serveraddress" bson:"serveraddress"`
}

type AlertMssg struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Status    string             `json:"status" bson:"status"`
	DateTIme  string             `json:"datetime" bson:"datetime"`
	People    int                `json:"numbers" bson:"numbers"`
	Hospital  Point              `json:"hospital" bson:"hospital"`
	VehicleID string             `json:"vehicle_id" bson:"vehicle_id"`
	Loc       []float64          `json:"accidentlocation" bson:"accidentlocation"`
	Sensors   Sensor             `json:"sensor" bson:"sensor"`
}
