package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/mch735/education/work3/api/middlewares"
	"github.com/mch735/education/work3/api/router"
	"github.com/mch735/education/work3/api/server"
	"github.com/mch735/education/work3/internal/config"
	"github.com/mch735/education/work3/internal/logger"
	"github.com/mch735/education/work3/internal/util"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello =)")
}

func main() {
	conf, err := config.Load()
	if err != nil {
		util.Fatal(err)
	}

	logger, err := logger.NewLogger(conf.LoggerConfig)
	if err != nil {
		util.Fatal(err)
	}

	router := router.NewRouter()
	router.Middleware(middlewares.ResultHandler{})
	router.Middleware(middlewares.StatusHandler{
		ProcessingFunc: func(method, path string, statusCode int, statusText string) {
			attributes := []any{
				slog.String("method", method),
				slog.String("path", path),
				slog.String("status", fmt.Sprintf("%d %s", statusCode, statusText)),
			}

			//nolint:mnd
			if statusCode >= 400 {
				logger.Error("request processed", attributes...)
			} else {
				logger.Info("request processed", attributes...)
			}
		},
	})
	router.HandleFunc("/", home)

	app, err := server.NewServer(conf.ServerConfig)
	if err != nil {
		util.Fatal(err)
	}

	app.Router(router)
	app.Run()
}
