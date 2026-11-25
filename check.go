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

type job struct {
	index  int
	target Target
}

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

const workers = 10

func checkAll(targets []Target) []Result {
	results := make([]Result, len(targets))

	jobs := make(chan job)

	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for j := range jobs {
				results[j.index] = check(j.target)
			}
		})
	}

	for i, t := range targets {
		jobs <- job{index: i, target: t}
	}

	close(jobs)

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
