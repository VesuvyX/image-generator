package main

import (
	"flag"
	"log"

	"github.com/VesuvyX/2-image-generator-server/configs"
	"github.com/VesuvyX/2-image-generator-server/internal/server"
)

var confPath = flag.String("conf-path", "../configs/.env", "Path to config env.")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}
