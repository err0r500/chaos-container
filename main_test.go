package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHappy(t *testing.T) {
	res := server(500)

	if res.StatusCode != 500 {
		log.Fatal("wrong statusCode returned")
	}
}

func server(input int) *http.Response {
	ts := httptest.NewServer(
		http.HandlerFunc(
			Handler(input),
		),
	)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
