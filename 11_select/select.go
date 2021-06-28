package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second


func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out after waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
