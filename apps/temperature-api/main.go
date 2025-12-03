package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Location string  `json:"location"`
	SensorID string  `json:"sensorId"`
	Value    float64 `json:"value"`
}

func main() {
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
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
		temp := rand.Float64()*10 + 20 // 20–30 градусов

		resp := TemperatureResponse{
			Location: location,
			SensorID: sensorID,
			Value:    temp,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8081", nil)
}
