package config

import (
	"os"
)

// ServerConfigInterface server config interface
type ServerConfigInterface interface {
	Port() string
	Mode() string
	FileServiceEndPoint() string
	AccountServiceEndPoint() string
	ProfileServiceEndPoint() string
}

// Server server config struct
type Server struct {
	port                   string
	mode                   string
	fileServiceEndPoint    string
	accountServiceEndPoint string
	profileServiceEndPoint string
}

// NewServerConfig create server config struct instance
func NewServerConfig() *Server {
	port := "5000"
	mode := "debug"
	fileServiceEndPoint := "http://localhost:5000/files"
	accountServiceEndPoint := "http://localhost:5000/accounts"
	profileServiceEndPoint := "http://localhost:5000/profiles"

	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("MODE"); env != "" {
		mode = env
	}
	if env := os.Getenv("FILE_API_ADDRESS"); env != "" {
		fileServiceEndPoint = env
	}
	if env := os.Getenv("ACCOUNT_API_ADDRESS"); env != "" {
		accountServiceEndPoint = env
	}
	if env := os.Getenv("PROFILE_API_ADDRESS"); env != "" {
		profileServiceEndPoint = env
	}
	server := &Server{
		port:                   port,
		mode:                   mode,
		fileServiceEndPoint:    fileServiceEndPoint,
		accountServiceEndPoint: accountServiceEndPoint,
		profileServiceEndPoint: profileServiceEndPoint,
	}
	if server.mode != "release" && server.mode != "debug" {
		panic("Unavailable gin mode")
	}
	return server
}

// Port get server port number
func (server *Server) Port() string {
	return server.port
}

// Mode get server mode
func (server *Server) Mode() string {
	return server.mode
}

// FileServiceEndPoint get file service endpoint
func (server *Server) FileServiceEndPoint() string {
	return server.fileServiceEndPoint
}

// AccountServiceEndPoint get account service endpoint
func (server *Server) AccountServiceEndPoint() string {
	return server.accountServiceEndPoint
}

// ProfileServiceEndPoint get profile service endpoint
func (server *Server) ProfileServiceEndPoint() string {
	return server.profileServiceEndPoint
}
