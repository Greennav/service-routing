package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

var mockupResponse string

type Location struct {
	Lat 	float64 	`json:"latitude"`
	Long	float64 	`json:"longitude"`
	Turn 	int64 		`json:"turn"`
}

type MockDataType struct {
	Name 		string 		`json:"algorithm"`
	BatteryStatus 	float64		`json:"batteryStatus"`
	Locations 	[]Location 	`json:"route"`
	Time 		float64 	`json:"time"`
	Distance 	float64 	`json:"distance"`
}

func init() {
	fileContent, _ := ioutil.ReadFile("mockup-response.json")
	var data MockDataType
	err := json.Unmarshal(fileContent, &data)
    	if err != nil {
        	log.Fatal(err)
    	}
	mockupResponse = string(fileContent)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, mockupResponse)
}

func main() {
	fmt.Println("Listening on http://localhost:6833. Ctrl+C to exit")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":6833", nil)
}
