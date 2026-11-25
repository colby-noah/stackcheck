package main

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
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
