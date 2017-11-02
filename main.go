package main

import (
	"fmt"
	"gol/models"
	"net/http"
)

func main() {
	http.HandleFunc("/", AuthMock)
	http.ListenAndServe(":3000", nil)
}

func AuthMock(w http.ResponseWriter, r *http.Request) {

  if r.Method != "GET" {
    w.Header().Set("Allow", "GET")
    http.Error(w, http.StatusText(405), 405)
    return
  }

	id := r.URL.Query().Get("user_key")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	auth, _ := models.Mock(id)

	// Write the album details as plain text to the client.
	fmt.Fprintf(w, "200")
}
