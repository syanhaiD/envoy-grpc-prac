package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":25000", nil)
}
