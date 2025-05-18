package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("running")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
