package app

import (
	"log"
	"lvl1/todo_http_server_with_json_storage/internal/middleware"
	"lvl1/todo_http_server_with_json_storage/internal/todo"
	"net/http"
	"strconv"
)

type Server struct {
	fullAddress string
}

func CreateServer(address string, port int) *Server {
	return &Server{fullAddress: address + ":" + strconv.Itoa(port)}

}
func (s Server) Start() {
	router := http.NewServeMux()
	todo.SetupToDoRouter(router)
	loggedRouter := middleware.LoggingMiddleware(router)
	log.Print("Listening on ", s.fullAddress)
	log.Fatal(http.ListenAndServe(s.fullAddress, loggedRouter))
}
