package todo

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func SetupToDoRouter(r *http.ServeMux) {
	r.HandleFunc("/echo", EchoRequest)
	r.HandleFunc("/createTask", CreateTask)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func EchoRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err)
		return
	}
	var rStr = string(rBody)
	log.Print("Getted request ", r.Method, " on ", r.URL.Path, " with content ", rStr)
	io.Copy(w, bytes.NewBuffer(rBody))
}
