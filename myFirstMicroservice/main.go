package main

import (
	"fmt"
	"net/http"
	"log"
)

const message = "hello, world\n"

func main() {
	fmt.Printf(message)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})
	err := http.ListenAndServe(":9400", mux)
	if err != nil {
		log.Fatalf("Server Failed: %", err)
	}
}