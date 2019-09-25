package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer races the ping of two web sites with a default timeout of 10 seconds
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

// ConfigurableRacer allows setting of a timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
