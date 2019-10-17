package main

import (
	"net/http"
	"os"
	"strconv"
)

const defaultStatusCode = 500

func Handler(statusCode int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(statusCode)
	}
}

func main() {
	http.HandleFunc("/", Handler(statusCodeToReturn()))
	http.ListenAndServe(":8080", nil)
}

func statusCodeToReturn() int {
	statusCode := defaultStatusCode

	if statusFromEnv, ok := os.LookupEnv("STATUS"); ok {
		if n, err := strconv.Atoi(statusFromEnv); err == nil {
			statusCode = n
		}
	}

	return statusCode
}
