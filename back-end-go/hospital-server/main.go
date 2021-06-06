package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/emergency", http.HandlerFunc(emergency))

	http.ListenAndServe(":8090", nil)
}

func emergency(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Analyzing Report...")
	data, _ := ioutil.ReadAll(req.Body)

	//print response  body
	fmt.Printf("%s\n", data)
	DispatchAmbulence()
	fmt.Fprint(w, "Ambulence Dispatched")
}

func DispatchAmbulence() {
	fmt.Println("Ambulence Dispatched")

}
