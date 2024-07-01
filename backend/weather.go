package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeatherData(lat, long string) (WeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&long=%s&appid=%s&units=metric", lat, long, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weatherData WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weatherData, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	if lat == "" || long == "" {
		http.Error(w, "lat and long parameters are required", http.StatusBadRequest)
		return
	}

	weatherData, err := getWeatherData(lat, long)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var temperatureStatus string
	if weatherData.Main.Temp < 25 {
		temperatureStatus = "cold"
	} else if weatherData.Main.Temp < 50 {
		temperatureStatus = "moderate"
	} else {
		temperatureStatus = "hot"
	}

	response := map[string]string{
		"condition": weatherData.Weather[0].Main,
		"temperature": temperatureStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}