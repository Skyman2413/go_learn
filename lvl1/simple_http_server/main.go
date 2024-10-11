package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		rBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %s", err)
			return
		}
		var rStr = string(rBody)
		log.Print("Getted request ", r.Method, " on ", r.URL.Path, " with content ", rStr)
		io.Copy(w, bytes.NewBuffer(rBody))
	})
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
