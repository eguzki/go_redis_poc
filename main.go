package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", AuthMock)
	http.ListenAndServe(":3000", nil)
}

func AuthMock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}
