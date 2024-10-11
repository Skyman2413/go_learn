package database

import (
	"encoding/json"
	"log"
	"lvl1/todo_http_server_with_json_storage/internal/model"
	"os"
)

type JsonDB struct {
	path string
	Jobs []model.Job `json:"jobs"`
}

func NewJsonDB(path string) *JsonDB {
	return &JsonDB{path: path, Jobs: make([]model.Job, 0)}
}

func (db JsonDB) Read() {
	err := json.Unmarshal([]byte(db.path), &db.Jobs)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func (db JsonDB) Write() {
	js, err := json.Marshal(db.Jobs)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(db.path, js, 0644)
}
func (db JsonDB) Create() {}
func (db JsonDB) Remove() {}
func (db JsonDB) Update() {}
func (db JsonDB) Delete() {}
