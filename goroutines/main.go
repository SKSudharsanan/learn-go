package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// WeatherData stores the parts of the response we care about
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Error bool
}

// fetchWeather fetches the weather for a given city and sends the data to a channel
func fetchWeather(city string, apiKey string, wg *sync.WaitGroup, c chan<- WeatherData) {
	defer wg.Done()
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		c <- WeatherData{Error: true}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data:", err)
		c <- WeatherData{Error: true}
		return
	}

	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing data:", err)
		c <- WeatherData{Error: true}
		return
	}

	c <- data
}

func main() {
	apiKey := "337a6fa61b76a8132b1b56ac3aa17eaf"
	cities := []string{"London", "Berlin", "New York", "chennai", "delhi", "udhuyeagdauyid"}

	c := make(chan WeatherData)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, apiKey, &wg, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	// Collect results with a timeout for the whole operation
	timeout := time.After(10 * time.Second)
	for {
		select {
		case data, ok := <-c:
			if !ok {
				return // Channel was closed and drained
			}
			if data.Error {
				fmt.Println("Failed to fetch data for:", data.Name)
			} else {
				fmt.Printf("The temperature in %s is %.2f°C\n", data.Name, data.Main.Temp)
			}
		case <-timeout:
			fmt.Println("Timed out waiting for all responses")
			return
		}
	}
}
