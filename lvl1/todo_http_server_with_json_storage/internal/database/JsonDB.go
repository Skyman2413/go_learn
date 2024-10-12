package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"lvl1/todo_http_server_with_json_storage/internal/model"
	"os"
	"reflect"
)

type JsonDB struct {
	path string
	Jobs map[string]model.Job `json:"jobs"`
}

var jsonDB *JsonDB = nil

func GetJsonDB(path string) *JsonDB {
	if jsonDB == nil {
		return &JsonDB{path: path, Jobs: make(map[string]model.Job)}
	}
	return jsonDB
}

func (db *JsonDB) Load() {
	err := json.Unmarshal([]byte(db.path), &db.Jobs)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func (db *JsonDB) Write() {
	js, err := json.Marshal(db.Jobs)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(db.path, js, 0644)
}

func (db *JsonDB) Create(Job model.Job) {
	db.Jobs[uuid.New().String()] = Job
}

func (db *JsonDB) Remove(Uuid string) {
	delete(db.Jobs, Uuid)
}

func (db *JsonDB) Update(Uuid string, Job model.Job) {
	db.Jobs[Uuid] = Job
}

func (db *JsonDB) Read(Uuid string) model.Job {
	return db.Jobs[Uuid]
}

func (db *JsonDB) Filter(key string, value string) map[string]model.Job {
	result := make(map[string]model.Job)

	for id, job := range db.Jobs {
		v := reflect.ValueOf(job)
		field := v.FieldByName(key)

		if field.IsValid() && field.Interface() == value {
			result[id] = job
		}
	}

	return result
}
