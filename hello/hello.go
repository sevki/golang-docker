package main

import (
	"os"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]bytes(os.Environ()))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
