package server

import "github.com/caarlos0/env"

// Server server config struct
type Server struct {
	Port string `env:"PORT" envDefault:"5000"`
	Mode string `env:"MODE" envDefault:"debug"`
}

// NewServer create server config struct instance
func NewServer() *Server {
	server := &Server{}
	env.Parse(server)
	if server.Mode != "release" && server.Mode != "debug" {
		panic("Unavailable gin mode")
	}
	return server
}
