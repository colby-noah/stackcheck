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

type Result struct {
	Name       string
	Up         bool
	StatusCode int
	Err        error
	Duration   time.Duration
}

func (r Result) Status() string {
	if r.Up {
		return "UP"
	}

	return "DOWN"
}

func check(target Target) Result {
	start := time.Now()

	resp, err := httpClient.Get(target.URL)
	elapsed := time.Since(start)

	if err != nil {
		return Result{Name: target.Name, Err: err, Duration: elapsed}
	}

	_ = resp.Body.Close()

	return Result{
		Name:       target.Name,
		Up:         resp.StatusCode < 400,
		StatusCode: resp.StatusCode,
		Duration:   elapsed,
	}
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
		r := check(target)

		status := r.Status()

		if r.Err != nil {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.Err, r.Duration)
		} else {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.StatusCode, r.Duration)
		}
	}
}
