package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	cfg, err := loadConfig("targets.yaml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	results := checkAll(cfg.Targets)

	for _, r := range results {
		status := r.Status()

		if r.Err != nil {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.Err, r.Duration.Round(time.Millisecond))
		} else {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.StatusCode, r.Duration.Round(time.Millisecond))
		}
	}
}
