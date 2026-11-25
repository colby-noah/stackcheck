package main

import (
	"net/http"
	"sync"
	"time"
)

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

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func checkAll(targets []Target) []Result {
	results := make([]Result, len(targets))

	var wg sync.WaitGroup
	for i, t := range targets {
		wg.Go(func() {
			results[i] = check(t)
		})
	}

	wg.Wait()
	return results
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
