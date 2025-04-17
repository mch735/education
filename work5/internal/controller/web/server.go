package web

import (
	"fmt"
	"net/http"

	"github.com/mch735/education/work5/config"
)

func NewServer(conf *config.HTTP) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	}
}
