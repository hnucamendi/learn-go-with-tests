package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speeds of servers", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 20)
		fastServer := makeDelayedServer(time.Millisecond * 0)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got: %q want: %v", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		delayedServer := makeDelayedServer(time.Millisecond * 25)

		defer delayedServer.Close()

		_, err := ConfigurableRacer(delayedServer.URL, delayedServer.URL, time.Millisecond*20)
		if err == nil {
			t.Errorf("expected an error but didint get one")
		}
	})

}

func makeDelayedServer(td time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(td)
		w.WriteHeader(http.StatusOK)
	}))

}
