package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cfg, err := loadConfig("targets.yaml")
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	results := checkAll(cfg.Targets)

	anyDown := false
	for _, r := range results {
		d := r.Duration.Round(time.Millisecond)
		status := r.Status()

		if !r.Up {
			anyDown = true
		}

		if r.Err != nil {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.Err, d)
		} else {
			fmt.Printf("%-15s %-6s %v %v\n", r.Name, status, r.StatusCode, d)
		}
	}

	if anyDown {
		os.Exit(1)
	}
}
