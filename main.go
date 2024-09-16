package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lat := os.Getenv("LAT")
	lon := os.Getenv("LON")
	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=%v&lon=%v", lat, lon)

	req, err := http.NewRequest("GET", url, http.NoBody)
	req.Header.Set("User-Agent", "github.com/sonlar/yr-melding")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
