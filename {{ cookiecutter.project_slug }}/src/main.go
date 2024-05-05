package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
