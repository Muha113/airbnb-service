package main

import (
	"net/http"

	"github.com/Muha113/airbnb-service/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.NewServer("./configs/server_config.json")
	if err != nil {
		logrus.Fatal(err)
	}

	addr := srv.Cfg.Host + ":" + srv.Cfg.Port

	logrus.Info("Starting server on -> ", addr)
	if err = http.ListenAndServe(addr, srv.Router); err != nil {
		logrus.Fatal(err)
	}
}
