package main

import (
	"WB-L2/develop/dev11/internal/config"
	"WB-L2/develop/dev11/internal/server"
	"log"
)

func main() {
	srvConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.NewServer()
	if err := srv.Run(srvConfig.Addr); err != nil {
		log.Fatal(err)
	}
}
