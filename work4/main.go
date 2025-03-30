package main

import (
	"encoding/csv"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mch735/education/work4/internal/scraper"
)

func main() {
	var conf scraper.Config

	err := cleanenv.ReadEnv(&conf)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	input := make(chan string)
	go func() {
		for _, arg := range os.Args[1:] {
			input <- arg
		}

		defer close(input)
	}()

	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"date", "url", "status_code", "title", "description"})
	if err != nil {
		panic(err)
	}

	output := scraper.Run(&conf, input)

	for data := range output {
		err := writer.Write(data.Fields())
		if err != nil {
			panic(err)
		}
	}
}
