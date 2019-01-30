package main

import (
	"net/http"

	"github.com/RequestsAllowedService/service"
)

func main() {

	http.HandleFunc("/requestAllowed", func(w http.ResponseWriter, r *http.Request) {
		service.RequestAllowed(w, r)
	})

	http.ListenAndServe(":80", nil)
}
