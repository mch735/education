package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mch735/education/work2/internal/storages/memory"
	"github.com/mch735/education/work2/internal/user"
)

var service = user.NewService(memory.NewUserRepo()) //nolint:gochecknoglobals

func main() {
	for {
		fmt.Println("Available commands: create, get, remove, list, filter, help, exit...")
		fmt.Print("> ")

		buf := bufio.NewReader(os.Stdin)

		command, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}

		args := strings.Fields(command)
		if len(args) > 0 {
			procesing(args)
		}
	}
}

func procesing(args []string) {
	switch args[0] {
	case "create":
		create(args[1:])
	case "get":
		get(args[1:])
	case "remove":
		remove(args[1:])
	case "list":
		list()
	case "filter":
		filter(args[1:])
	case "help":
		help()
	case "exit":
		exit()
	default:
		inputError()
	}
}

func create(args []string) {
	cmd := flag.NewFlagSet("create", flag.ContinueOnError)
	name := cmd.String("name", "", "user name")
	email := cmd.String("email", "", "user email")
	role := cmd.String("role", "", "user role (admin, user, gues)\n")

	if err := cmd.Parse(args); err != nil {
		return
	}

	record, err := service.CreateUser(*name, *email, *role)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(record)
}

func get(args []string) {
	cmd := flag.NewFlagSet("get", flag.ContinueOnError)
	id := cmd.String("id", "", "user id")

	if err := cmd.Parse(args); err != nil {
		return
	}

	record, err := service.GetUser(*id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(record)
}

func remove(args []string) {
	cmd := flag.NewFlagSet("remove", flag.ContinueOnError)
	id := cmd.String("id", "", "user id")

	if err := cmd.Parse(args); err != nil {
		return
	}

	err := service.RemoveUser(*id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("User removed...")
}

func list() {
	records := service.ListUsers()

	if len(records) > 0 {
		for _, record := range records {
			fmt.Println(record)
		}
	} else {
		fmt.Println("Users not found...")
	}
}

func filter(args []string) {
	cmd := flag.NewFlagSet("filter", flag.ContinueOnError)
	role := cmd.String("role", "", "user role (admin, user, gues)")

	if err := cmd.Parse(args); err != nil {
		return
	}

	records := service.ListUsersWithRole(*role)
	if len(records) > 0 {
		for _, record := range records {
			fmt.Println(record)
		}
	} else {
		fmt.Println("Users not found...")
	}
}

func exit() {
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
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
	fmt.Println("      example: get -id=f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	fmt.Println("      params:")
	fmt.Println("        id  - user id")
	fmt.Println()
	fmt.Println("  remove - delete user by id")
	fmt.Println("      example: remove f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	fmt.Println("      params:")
	fmt.Println("        id  - user id")
	fmt.Println()
	fmt.Println("  list - list users")
	fmt.Println("      example: list")
	fmt.Println()
	fmt.Println("  filter - filter user by role")
	fmt.Println("      example: filter -role=admin")
	fmt.Println("      params:")
	fmt.Println("        role - one of 'admin', 'user' or 'guest'")
	fmt.Println()
	fmt.Println("  exit - quit")
	fmt.Println()
}

func inputError() {
	fmt.Println("expected 'create', 'get', 'remove', 'list' or 'filter' subcommands")
	fmt.Println()
}
