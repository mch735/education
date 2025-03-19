package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mch735/education/work3/api/router"
	"github.com/mch735/education/work3/internal/config"
	"github.com/mch735/education/work3/internal/util"
)

const (
	ReadTimeout  time.Duration = 5 * time.Second
	WriteTimeout time.Duration = 3 * time.Second
)

type Server struct {
	http.Server
}

func NewServer(conf config.ServerConfig) (*Server, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("invalid server settings: %w", err)
	}

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return &Server{
		http.Server{
			Addr:         addr,
			ReadTimeout:  ReadTimeout,
			WriteTimeout: WriteTimeout,
		},
	}, nil
}

func (s *Server) Router(router *router.Router) {
	s.Server.Handler = router
}

func (s *Server) Run() {
	err := s.Server.ListenAndServe()
	if err != nil {
		util.Fatal(err)
	}
}
