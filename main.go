package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	Add  string = "add"
	List string = "list"
	Help string = "help"
	Auth string = "auth"
)

func auth(username string, password string, baseUrl string) {
	// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
	body, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(baseUrl+"/jwt/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

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
			auth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("BASE_URL"))
		case Help:
			help()
		case "":
			help()
		default:
			help()
		}
	}
}
