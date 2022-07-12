package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	if b, err := json.Marshal(req.Header); err == nil {
		fmt.Printf("{\"Headers\": %v}\n", string(b))
	}
	if b, err := io.ReadAll(req.Body); err == nil {
		fmt.Println(string(b))
	}

	w.Header().Set("x-response-from", "echo")
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	fmt.Fprintf(w, "")
}

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8090", nil)
}
