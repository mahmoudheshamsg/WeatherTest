package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

	}
	fmt.Println(os.Getenv("OPENWEATHERMAPAPIKEY"))
}
