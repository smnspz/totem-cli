package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/smnspz/totem-cli/auth"
)

const (
	Add  string = "add"
	List string = "list"
	Help string = "help"
	Auth string = "auth"
)

func add() {
	fmt.Println("Adding a new entry")
}

func list() {
	fmt.Println("Listing all the entries")
}

func help() {
	fmt.Println("Printing help")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	for _, arg := range args {

		switch arg {
		case Add:
			add()
		case List:
			list()
		case Auth:
			auth.GetToken(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("BASE_URL"))
		case Help:
			help()
		case "":
			help()
		default:
			help()
		}
	}
}
