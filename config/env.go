package config

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)
var (
	CSTZone = time.FixedZone("CST",8*3600)
)
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}