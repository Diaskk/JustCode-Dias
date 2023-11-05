package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/requestB", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/data")
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		fmt.Fprint(w, "Data from Service 2: ", resp.Status)
	})

	http.ListenAndServe(":8080", nil)
}
