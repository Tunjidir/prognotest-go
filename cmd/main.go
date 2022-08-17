//  Copyright © 2022 Olatunji Oniyide <olatunji4you@gmail.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	username           string = "prognotest"
	DefaultContentType string = "application/json; charset=utf-8"
	baseUrl            string = "http://api.geonames.org/findNearByWeatherJSON"
)

var client *http.Client

// Weather is the weather object that will
type Weather struct {
	WeatherObservation *WeatherObservation `json:"weatherObservation"`
}

// WeatherObservation holds specific details as regards to the weather
type WeatherObservation struct {
	WeatherCondition     string  `json:"weatherCondition"`
	Clouds               string  `json:"clouds"`
	Observation          string  `json:"observation"`
	WindDirection        float64 `json:"windDirection"`
	ICAO                 string  `json:"ICAO"`
	Elevation            float64 `json:"elevation"`
	CountryCode          string  `json:"countryCode"`
	CloudsCode           string  `json:"cloudsCode"`
	Lng                  float64 `json:"lng"`
	DewPoint             string  `json:"dewPoint"`
	Temperature          string  `json:"temperature"`
	WindSpeed            string  `json:"windSpeed"`
	Humidity             float64 `json:"humidity"`
	Datetime             string  `json:"datetime"`
	StationName          string  `json:"stationName"`
	Lat                  float64 `json:"lat"`
	HectoPascAltimeter   float64 `json:"hectoPascAltimeter"`
	WeatherConditionCode string  `json:"weatherConditionCode"`
}

func init() {
	client = &http.Client{}
	http.HandleFunc("/api", WeatherHandler)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", DefaultContentType)
	lat := r.URL.Query().Get("latitude")
	lng := r.URL.Query().Get("longitude")
	url := baseUrl + "?lat=" + lat + "&lng=" + lng + "&username=" + username
	weather := &Weather{
		WeatherObservation: &WeatherObservation{},
	}
	bt := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, bt)
	req.Header.Set("Content-Type", DefaultContentType)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Errorf("error handling client request %s", err)
	}
	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, weather)
	fmt.Println(weather.WeatherObservation.Humidity)
	fmt.Println(weather.WeatherObservation.ICAO)
}

func main() {
	log.Fatal(http.ListenAndServe(":9080", nil))
}
