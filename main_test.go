package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHappy(t *testing.T) {
	res := server(500, 100)

	if res.StatusCode != 500 {
		log.Fatal("wrong statusCode returned")
	}
}

func TestHappyDisabled(t *testing.T) {
	res := server(500, 0)

	if res.StatusCode != 200 {
		log.Fatal("wrong statusCode returned")
	}
}

func server(input, freq int) *http.Response {
	ts := httptest.NewServer(
		http.HandlerFunc(
			Handler(input, freq),
		),
	)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
