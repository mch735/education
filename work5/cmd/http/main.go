package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "create":
			create(ctx, www, os.Args[2:])
		case "get":
			get(ctx, www, os.Args[2])
		case "del":
			del(ctx, www, os.Args[2])
		case "list":
			list(ctx, www)
		case "help":
			help()
		default:
			inputError()
		}
	} else {
		inputError()
	}
}

func create(ctx context.Context, service *web.Client, options []string) {
	cmd := flag.NewFlagSet("create", flag.ContinueOnError)
	name := cmd.String("name", "", "user name")
	email := cmd.String("email", "", "user email")
	role := cmd.String("role", "", "user role (admin, user, gues)")

	err := cmd.Parse(options)
	if err != nil {
		return
	}

	user, err := service.Create(ctx, *name, *email, *role)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*user)
	}
}

func get(ctx context.Context, service *web.Client, id string) {
	user, err := service.GetUserByID(ctx, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*user)
	}
}

func del(ctx context.Context, service *web.Client, id string) {
	err := service.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User deleted...")
	}
}

func list(ctx context.Context, service *web.Client) {
	users, err := service.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}

	if len(users) > 0 {
		for _, user := range users {
			fmt.Println(*user)
		}
	} else {
		fmt.Println("Users not found...")
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("  create - create user")
	fmt.Println("      example: create -name=jon -email=1@1.com -role=user")
	fmt.Println("      params:")
	fmt.Println("        name  - username")
	fmt.Println("        email - user@emanple.com")
	fmt.Println("        role  - one of 'admin', 'user' or 'guest'")
	fmt.Println()
	fmt.Println("  get - find user by id")
	fmt.Println("      example: get f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	fmt.Println()
	fmt.Println("  del - delete user by id")
	fmt.Println("      example: remove f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	fmt.Println()
	fmt.Println("  list - list users")
	fmt.Println("      example: list")
	fmt.Println()
	fmt.Println("  help - information of subcommands")
	fmt.Println("      example: help")
	fmt.Println()
}

func inputError() {
	fmt.Println("expected 'create', 'get', 'del', 'list' or 'help' subcommands")
	fmt.Println()
}
