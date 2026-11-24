package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.yaml.in/yaml/v3"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

type Target struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Config struct {
	Targets []Target `yaml:"targets"`
}

func main() {
	data, err := os.ReadFile("./targets.yaml")
	if err != nil {
		log.Fatalf("reading config: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("reading config: %v", err)
	}

	for _, target := range config.Targets {
		start := time.Now()

		resp, err := httpClient.Get(target.URL)
		if err != nil {
			fmt.Printf("%-15s %-6s %v\n", target.Name, "DOWN", err)
			continue
		}

		resp.Body.Close()
		fmt.Printf("%-15s %-6s %d %s\n", target.Name, "UP", resp.StatusCode, time.Since(start))
	}
}
