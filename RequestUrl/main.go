package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "req.URL: %v\n", req.URL)
	})

	http.ListenAndServe(":8080", nil)
}