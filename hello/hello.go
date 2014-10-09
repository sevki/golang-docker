package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		for _, v := range os.Environ() {

			w.Write([]byte(fmt.Sprintf("%s\n", v)))
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
