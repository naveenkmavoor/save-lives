package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"encoding/json"
	"log"
	"os"
	"time"

	"Rest_api/models"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSensorVal(mongoclient *mongo.Client) {

	var flag bool = false
	alertmssg := models.AlertMssg{}

	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("mqtt://tailor.cloudmqtt.com:12189")
	opts.SetUsername("user name")
	opts.SetPassword("password")
	opts.SetClientID("rasp-pi-go")
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println("SetDefaultPublishHandler : ", msg.Topic(), string(msg.Payload()))
	})

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("some_topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//define a function for the default message handler
	var callback MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Listening to Topic: %s\n", msg.Topic())
		fmt.Printf("Recieved Payload Message : %s\n", msg.Payload())
		fmt.Println("Analyzing data...")
		err := json.Unmarshal([]byte(msg.Payload()), &alertmssg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Analyzing Severity ...")
		fmt.Println("Determined as ")
		//determine the severe of injury
		severity(&alertmssg)

		//determine nearby hospital
		nearbyHospitals := nearBy(&alertmssg, mongoclient)

		fmt.Println("Forwarding Alert Report to Nearby Hospital.....")

		//send file report to hospital
		sendReport(&alertmssg, nearbyHospitals)

		//adding the values to the database
		insertDB(&alertmssg, mongoclient)
		return

	}

	//subscribe
	if token := c.Subscribe("mqtt", 0, callback); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	for flag == false {
		time.Sleep(1 * time.Second)
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("some_topic"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func insertDB(alertmssg *models.AlertMssg, mongoclient *mongo.Client) {
	alertmssg.Id = primitive.NewObjectID() //create an object id
	collection := mongoclient.Database("fir").Collection("cases")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// store the case in mongodb
	_, insertErr := collection.InsertOne(ctx, alertmssg)
	if insertErr != nil {
		fmt.Println("Insertone Mongo ERROR", insertErr)
		os.Exit(1)
	}
	fmt.Println("Case Result added to database")

}

func severity(alertmessage *models.AlertMssg) {
	sensorReadings := alertmessage.Sensors
	accelerometerData := sensorReadings.Accelerometer
	soundData := sensorReadings.Sound
	s := accG(accelerometerData, soundData)
	alertmessage.Status = s
	fmt.Println(s)
}

func nearBy(alertmssg *models.AlertMssg, mongoclient *mongo.Client) []models.Point {
	// Location is a GeoJSON type.
	nearbyHospitals := []models.Point{}
	c := alertmssg.Loc
	latitude := c[1]
	longitude := c[0]
	collection := mongoclient.Database("hospitals").Collection("place")

	location := NewPoint(longitude, latitude)
	filter := bson.D{
		{"location",
			bson.D{
				{"$near", bson.D{
					{"$geometry", location},
					{"$maxDistance", 10000},
				}},
			}},
		{"status", "available"},
	}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {

		fmt.Println("Error occured at nearby")

	}
	for cur.Next(context.TODO()) {
		var p models.Point
		err := cur.Decode(&p)
		if err != nil {
			fmt.Println("Could not decode Point")

		}
		nearbyHospitals = append(nearbyHospitals, p)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return nearbyHospitals

}

func accG(accelerometerData []float64, sound int) string {
	xval := accelerometerData[0]
	yval := accelerometerData[1]
	zval := accelerometerData[2]

	accg := math.Sqrt(xval*xval + yval*yval + zval*zval) // calculate the magnitude of impact

	if accg > 40.0 || sound >= 180 {
		return "Severe Accident"

	} else if accg > 20.0 {
		return "Medium Accident"

	} else if accg > 4.0 {
		return "Mild Accident"
	}
	return "No Accident"
}

// NewPoint returns a GeoJSON Point with longitude and latitude.
func NewPoint(long, lat float64) models.Location {
	return models.Location{
		Type:        "Point",
		Coordinates: []float64{long, lat},
	}
}
func sendReport(alertmssg *models.AlertMssg, nearbyHospitals []models.Point) {

	if len(nearbyHospitals) == 0 {
		return
	}

	fmt.Println("NEARBY HOSPITALS DETECTED! ")

	//finding all possiblity to successfully connect a nearby hospital
	for _, val := range nearbyHospitals {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(val)
		if err != nil {
			return
		}
		response, err := http.Post(val.Endpoint, "application/json", b)
		if err != nil {
			fmt.Println("Failed to connect the hospital....")
			//on failing would lead to search next nearby hospital
			continue
		}

		fmt.Printf("Available hospital : %v\n", val)
		alertmssg.Hospital = val
		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		//print response body
		fmt.Printf("%s\n", data)
		fmt.Println("Acknowledgement from Hospital")
		//on successful alert exit from the loop
		return 
	}

}
