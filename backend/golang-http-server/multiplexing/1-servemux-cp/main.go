package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah http.ServeMux yang memiliki dua route
// Route pertama yaitu "/time" yang menghandle TimeHandler
// Route kedua yaitu "/hello" yang menghandle SayHelloHandler

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: replace this
		t := time.Now()
		output := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		fmt.Fprint(w, output)
	}
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // TODO: replace this
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprint(w, "Hello there")
			return
		}
		output := fmt.Sprintf("Hello, %v!", name)
		fmt.Fprint(w, output)
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}
