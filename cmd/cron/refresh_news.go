package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

const (
	fileConfig = ".env"
)

func init() {
	// Initialize Viper
	viper.SetConfigFile(fileConfig)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Config file not found: %v", err)
	}
}

func main() {
	makeRequest()
}

func makeRequest() {
	url := "https://mceasy.mainhaustradeclub.com/mceasy/news?refresh=true"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	clientKey := viper.GetString("x.client.key")
	req.Header.Add("X-Client-Key", clientKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Refresh news request completed with status: %s", resp.Status)
}
