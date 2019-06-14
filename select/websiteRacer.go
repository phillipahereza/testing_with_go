package websiteracer

import (
	"time"
	"errors"
	"net/http"
)

// Racer returns the fastest website
func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", errors.New("timeout")
	}
	
}




func ping(url string) (chan bool) {
	ch := make(chan bool)
	go func(){
		http.Get(url)
		ch <- true
	}()
	return ch
}

