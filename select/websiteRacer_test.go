package websiteracer

import (
	"net/http"
	"time"
	"net/http/httptest"
	"testing"
)

func TestRacer(t *testing.T) {

	t.Run("normal run", func(t *testing.T){
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, _ := Racer(slowURL, fastURL, 2 * time.Second)
	
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
		
	})

	t.Run("returns an error if a server doesn't respond within 1s", func(t *testing.T){
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 1 * time.Second)
		if err == nil {
			t.Error("expected and error but did not get one")
		}

	})


}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}