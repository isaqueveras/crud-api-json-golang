package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá Isaque, diga, Olá mundo!")
	})

	http.ListenAndServe(":1337", nil)
}
