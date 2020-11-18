package server

import (
	"net/http"

	"github.com/Muha113/airbnb-service/pkg/config"
	"github.com/Muha113/airbnb-service/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Config : contains all needful information about server config
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// Server : contains all needful information about server
type Server struct {
	Cfg    *Config
	Logger *logrus.Logger
	Router *mux.Router
}

// NewServer : creates new server
func NewServer(serverCfg string) (*Server, error) {
	cfg := &Config{}

	err := config.NewConfig(serverCfg, cfg)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		Cfg:    cfg,
		Logger: logrus.New(),
		Router: mux.NewRouter(),
	}

	return srv, nil
}

//Start : start server
func (s *Server) Start(addr string) error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.Logger.Info("Starting server on -> " + addr)
	return http.ListenAndServe(addr, s.Router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}
	s.Logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.Router.HandleFunc("/register", handlers.Login()).Methods("POST")
}
