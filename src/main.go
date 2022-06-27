package main

import (
	"net/http"
	// pb "github.com/Ivanf1/go-protobuf/protos"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	http.ListenAndServe(":3000", nil)
}
