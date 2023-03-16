package main

import (
	"fmt"
	"os"
)

const (
	Add  string = "add"
	List string = "list"
	Help string = "help"
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

	for _, arg := range args {

		fmt.Println(arg)

		switch arg {
		case Add:
			add()
		case List:
			list()
		case Help:
			help()
		case "":
			help()
		default:
			help()
		}
	}
}
