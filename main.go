package main

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func main() {
	urls := []string{"https://colby-noah.github.io/asteroids/"}

	for _, u := range urls {
		start := time.Now()

		resp, err := httpClient.Get(u)
		if err != nil {
			fmt.Printf("%s DOWN: %v\n", u, err)
			continue
		}

		resp.Body.Close()
		fmt.Printf("%-45s %d %s\n", u, resp.StatusCode, time.Since(start))
	}
}
