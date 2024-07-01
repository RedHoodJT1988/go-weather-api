package main 

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherHandler(t *testing.T) {
	os.Setenv("OPENWEATHER_API_KEY", "test_api_key")

	req, err := http.NewRequest("GET", "/weather?lat=35&long=139", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"condition": 'Rain", "temperature":"moderate"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestWeatherHandlerMissingParmas(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"condition": "Rain", "temperature": "moderate"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestWeatherHandlerMissingParams(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
