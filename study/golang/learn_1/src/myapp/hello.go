package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello11", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
