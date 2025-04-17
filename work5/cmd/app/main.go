package main

import (
	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(conf)
}
