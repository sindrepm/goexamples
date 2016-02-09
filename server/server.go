package main

import (
	"fmt"
	"net/http"

	"github.com/sindrepm/goexamples/api"
)

func main() {
	http.HandleFunc("/version", api.Info)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
