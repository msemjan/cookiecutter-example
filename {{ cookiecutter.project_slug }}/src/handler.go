package main

import (
	"fmt"
	{%- if cookiecutter.logging -%}"log"{% endif %}
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	{%- if cookiecutter.logging -%}log.Println("/index endpoint was hit"){% endif %}
	fmt.Fprintf(w, fmt.Sprintf("Hello %s", "{{ cookiecutter.author }}"))
}
