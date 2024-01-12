package patterns

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type ServerOption func(server *Server)

func New(options ...ServerOption) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	// todo
	return nil
}

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) ServerOption {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func main() {
	svr := New(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
