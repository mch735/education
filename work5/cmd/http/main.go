package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mch735/education/work5/internal/controller/web"
)

//go:embed http.json
var s []byte

type config struct {
	URL string `json:"url"`
}

func main() {
	var conf config

	err := json.Unmarshal(s, &conf)
	if err != nil {
		panic(err)
	}

	www := web.NewClient(conf.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	users, err := www.GetUsers(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
