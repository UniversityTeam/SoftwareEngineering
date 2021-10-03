package main

import (
	"fmt"
	"net/http"
)

func feature() {
	fmt.Println("New feature delivered!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running on Go...")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Now time is:")
	})
	http.ListenAndServe(":8795", nil)
}