package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerInfo struct {
	ServerName string
	Version    string
	PoweredBy  string
}

func Info(w http.ResponseWriter, r *http.Request) {
	info := ServerInfo{ServerName: "GoHttp", Version: "0.1a", PoweredBy: "Golang"}
	jsonResponse, err := json.MarshalIndent(info, "", "   ")
	if err != nil {
		fmt.Printf("Error: %v", err)
		w.Write([]byte("An error occured!"))
		return
	}

	w.Write(jsonResponse)
}
