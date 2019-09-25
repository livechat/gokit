package web

import (
	"net/http"

	"github.com/sokool/gokit/web/server"
)

//func Server() {
//	server.NewRouter()
//	server.Run("addr", server.NewRouter())
//}

type Server struct{ *server.Router }

func NewServer() *Server { return &Server{server.NewRouter()} }

func (s *Server) Run(addr string) error { return http.ListenAndServe(addr, s) }
