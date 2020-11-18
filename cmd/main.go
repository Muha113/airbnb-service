package main

import (
	"log"

	"github.com/Muha113/airbnb-service/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.NewServer("./configs/server_config.json")
	if err != nil {
		logrus.Fatal(err)
	}
	addr := srv.Cfg.Host + ":" + srv.Cfg.Port
	if err = srv.Start(addr); err != nil {
		log.Fatal(err)
	}
}
