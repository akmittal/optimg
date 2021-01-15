package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/akmittal/optimg/server/pkg/router"
	"github.com/go-chi/chi"
)

type Server struct {
	srv *http.Server
}

func Get(addr string, rtr *router.Router) *Server {
	server := &Server{
		srv: &http.Server{},
	}
	server = server.WithAddr(addr)
	server = server.WithRouter(rtr.Mux)
	return server
}

func (s *Server) WithAddr(addr string) *Server {
	s.srv.Addr = addr
	return s
}

func (s *Server) WithErrLogger(l *log.Logger) *Server {
	s.srv.ErrorLog = l
	return s
}

func (s *Server) WithRouter(router *chi.Mux) *Server {
	s.srv.Handler = router
	return s
}

func (s *Server) Start() error {
	if len(s.srv.Addr) == 0 {
		return errors.New("Server missing address")
	}

	if s.srv.Handler == nil {
		return errors.New("Server missing handler")
	}

	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}
