package main

import "github.com/gyu-young-park/lck_data_api/api"

func main() {
	server := api.NewHTTPServer()
	server.StartServer()
}