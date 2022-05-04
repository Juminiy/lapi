package main

import (
	"lapi/config"
	_ "lapi/docs"
	"lapi/server"
	"log"
)
func main() {
	app := server.InitServer()
	log.Fatalln(app.Listen(config.Config("SERVER_PORT")))
}
