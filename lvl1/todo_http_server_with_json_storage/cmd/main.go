package main

import "lvl1/todo_http_server_with_json_storage/internal/app"

func main() {
	server := app.CreateServer("localhost", 8080)
	server.Start()
}
