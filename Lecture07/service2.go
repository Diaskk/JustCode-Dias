package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/requestA", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8080/data")
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		fmt.Fprint(w, "data from service 1: ", resp.Status)
	})

	http.ListenAndServe(":8081", nil)
}
