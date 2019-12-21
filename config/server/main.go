package server

import "github.com/caarlos0/env"

// Server server config struct
type Server struct {
	Port string `env:"PORT" envDefault:"5000"`
}

// NewServer create server config struct instance
func NewServer() *Server {
	server := &Server{}
	env.Parse(server)
	return server
}
