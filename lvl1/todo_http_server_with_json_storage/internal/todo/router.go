package todo

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"lvl1/todo_http_server_with_json_storage/internal/database"
	"lvl1/todo_http_server_with_json_storage/internal/model"
	"net/http"
)

var jsonDB = database.GetJsonDB("")

func SetupToDoRouter(r *http.ServeMux) {
	r.HandleFunc("/echo", EchoRequest)
	r.HandleFunc("/createTask", CreateTask)
	r.HandleFunc("/deleteTask", RemoveTask)
	r.HandleFunc("/updateTask", UpdateTask)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var j model.Job
	rBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &j)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	jsonDB.Create(j)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	type UUID struct {
		Id string `json:"uuid"`
	}
	var uuid UUID
	rBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &uuid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	jsonDB.Remove(uuid.Id)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	type JobUUID struct {
		Id  string    `json:"uuid"`
		Job model.Job `json:"job"`
	}
	var j JobUUID
	rBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &j)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	jsonDB.Update(j.Id, j.Job)
}

func FilterTasks(w http.ResponseWriter, r *http.Request) {
	// TODO think about filtering
}

func EchoRequest(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %s", err)
		return
	}
	var rStr = string(rBody)
	log.Print("Getted request ", r.Method, " on ", r.URL.Path, " with content ", rStr)
	io.Copy(w, bytes.NewBuffer(rBody))
}
