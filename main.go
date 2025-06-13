package main

import (
	"fmt"
	"net/http"
)

func main() {
	access_times := 0

	increment := make(chan bool)
	go func() {
		for {
			<-increment
			access_times++
		}
	}()

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		increment <- true
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Access times: %d\n", access_times)))
	}))
}
