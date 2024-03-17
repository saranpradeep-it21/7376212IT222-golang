package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/reg-no", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT222...")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Saranpradeep R...")
	})
	http.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Micro_Service using Go-Lang...")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/reg-no\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
