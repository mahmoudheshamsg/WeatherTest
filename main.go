package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func KelvinToCelsius(Kelvin float64) float64 {
	return Kelvin - 273.15
}
func GetWeather(city string) float64 {
	//JSON struct template to get the temperature from the API
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
	//Making an API call to get the weather with the city name and APIKEY
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, os.Getenv("OPENWEATHERMAPAPIKEY"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//Reading the entire JSON at once
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)

	}
	//decoding the JSON and extracting only the temp value
	json.Unmarshal(body, &data)
	return KelvinToCelsius(data.Main.Temp)
}
func main() {
	city := flag.String("city", "Cairo", "Name of the City")
	flag.Parse()
	//load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading dotenv file", err)
	}
	fmt.Printf("%0.1f°C\n", GetWeather(*city))
}
