package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		for _, v := range os.Environ() {

			w.Write([]byte(fmt.Sprintf("%s\n", v)))
		}
		w.Write([]byte("\n################################\n\n"))
		addrs, _ := net.InterfaceAddrs()
		for _, v := range addrs {

			w.Write([]byte(fmt.Sprintf("%s\n", v.String())))
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
