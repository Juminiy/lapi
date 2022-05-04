package main

import (
	"log"
	"zhaoxin-api/config"
	_ "zhaoxin-api/docs"
	"zhaoxin-api/server"
)
func main() {
	app := server.InitServer()
	log.Fatalln(app.Listen(config.Config("SERVER_PORT")))
}
