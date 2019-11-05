package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const defaultStatusCode = 500
const defaultProbability = 100

func Handler(statusCode, frequency int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		if rand.Intn(100) < frequency {
			w.WriteHeader(statusCode)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", Handler(statusCodeToReturn(), failurProbability()))
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

func failurProbability() int {
	frequency := defaultProbability

	if freqFromEnv, ok := os.LookupEnv("FREQ"); ok {
		if n, err := strconv.Atoi(freqFromEnv); err == nil {
			frequency = n
		}
	}

	return frequency
}
