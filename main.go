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
	fmt.Fprintf(w, "")
}

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8090", nil)
}
