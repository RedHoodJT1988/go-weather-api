package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type WeatherResponse struct {
    Weather []struct {
        Main string `json:"main"` // e.g., "Rain", "Snow"
    } `json:"weather"`
    Main struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := os.Getenv("OPENWEATHER_API_KEY")
    if apiKey == "" {
        http.Error(w, "Missing OpenWeather API key", http.StatusInternalServerError)
        return
    }

    lat := r.URL.Query().Get("lat")
    lon := r.URL.Query().Get("lon")

    if lat == "" || lon == "" {
        http.Error(w, "Missing latitude or longitude", http.StatusBadRequest)
        return
    }

    apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, apiKey)

    resp, err := http.Get(apiURL)
    if err != nil {
        http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var weatherData WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
        http.Error(w, "Error parsing weather data", http.StatusInternalServerError)
        return
    }

    condition := weatherData.Weather[0].Main
    var temperatureCategory string

    switch {
    case weatherData.Main.Temp < 10:
        temperatureCategory = "cold"
    case weatherData.Main.Temp < 25:
        temperatureCategory = "moderate"
    default:
        temperatureCategory = "hot"
    }

    result := fmt.Sprintf("Condition: %s, Temperature: %s", condition, temperatureCategory)
    w.Write([]byte(result))
}

func main() {
    http.HandleFunc("/weather", weatherHandler)
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}
