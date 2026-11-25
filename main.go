package main

import (
	"fmt"
	"log"
)

func main() {
	cfg, err := loadConfig("targets.yaml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	for _, target := range cfg.Targets {
		r := check(target)

		status := r.Status()

		if r.Err != nil {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.Err, r.Duration)
		} else {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.StatusCode, r.Duration)
		}
	}
}
