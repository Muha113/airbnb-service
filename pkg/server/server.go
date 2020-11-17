package server

import (
	"net/http"

	"github.com/Muha113/airbnb-service/pkg/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Server struct {
	Cfg    *Config
	Router *mux.Router
}

func NewServer(serverCfg string) (*Server, error) {
	cfg := &Config{}

	err := config.NewConfig(serverCfg, cfg)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		Cfg:    cfg,
		Router: mux.NewRouter(),
	}

	srv.Router.HandleFunc("/", srv.Stub).Methods(http.MethodGet)

	return srv, nil
}

func (s *Server) Stub(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
	logrus.Info("Sending response on \"/\"...")
}
