package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Location string  `json:"location"`
	SensorID string  `json:"sensorId"`
	Value    float64 `json:"value"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	rand.Seed(time.Now().UnixNano())
	temp := 18 + rand.Float64()*10

	resp := TemperatureResponse{
		Location: location,
		SensorID: sensorID,
		Value:    temp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/temperature", temperatureHandler)
	mux.HandleFunc("/temperature/", temperatureHandler)

	log.Println("temperature-api is running on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal("server error:", err)
	}
}