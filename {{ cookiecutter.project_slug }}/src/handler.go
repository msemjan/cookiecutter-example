package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/index endpoint was hit")
	fmt.Fprintf(w, fmt.Sprintf("Hello %s", "{{ cookiecutter.author }}"))
}
