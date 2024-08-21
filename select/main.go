package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout time.Duration = time.Second * 10

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func() {
		http.Get(url)
		close(c)
	}()
	return c
}

func Racer(a, b string) (winnder string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}

func main() {
	Racer("https://youtube.com", "https://redventures.com")
}
