package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

func feature() {
	fmt.Println("New feature delivered!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running on Go...")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now()
		response := fmt.Sprintf("Now time is: %d-%02d-%02dT%02d:%02d:%02d-00:00\n",
			tm.Year(), tm.Month(), tm.Day(),
			tm.Hour(), tm.Minute(), tm.Second())
		fmt.Fprintf(w, response)
		jsonData, err := json.Marshal(TimeS{Time: tm.Format(time.RFC3339)})
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))
		fmt.Fprintf(w, string(jsonData))
	})

	http.ListenAndServe(":8795", nil)
}
