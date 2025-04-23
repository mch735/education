package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	usergrpc "github.com/mch735/education/work5/internal/controller/grpc"
)

//go:embed grpc.json
var s []byte

type config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	var conf config

	err := json.Unmarshal(s, &conf)
	if err != nil {
		panic(err)
	}

	grpc, err := usergrpc.NewClient(conf.Host, conf.Port)
	if err != nil {
		panic(err)
	}
	defer grpc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	users, err := grpc.GetUsers(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
