package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		Routes().ServeHTTP(wr, req)
		// TODO: answer here
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/echo?data=Hi!", nil)
		Routes().ServeHTTP(wr, req)
		// TODO: answer here
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		a, _ := strconv.Atoi(r.URL.Query().Get("a"))
		b, _ := strconv.Atoi(r.URL.Query().Get("b"))
		fmt.Fprintf(w, "%d", a+b)

		// TODO: answer here
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Errorf("%s", r.URL.Query().Get("a"))
		// wr := httptest.NewRecorder()
		// req := httptest.NewRequest("GET", "/hellojson", nil)
		// server.Routes().ServeHTTP(wr, req)
		// TODO: answer here
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
