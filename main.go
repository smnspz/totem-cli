package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml"
	"github.com/smnspz/totem-cli/auth"
)

const (
	Add  string = "add"
	List string = "list"
	Help string = "help"
	Auth string = "auth"
)

type User struct {
	email    string
	password string
}

const totemConfigVar string = "TOTEM_CONFIG"
const totemConfig string = ".totemconfig"

func parseConfigs(configFile string) *User {
	var email, password string
	body, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}
	config, err := toml.Load(string(body))
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}
	email = config.Get("user.email").(string)
	password = config.Get("user.password").(string)
	return &User{email, password}
}

func setTotemConfigs(pathToConfig string) {
	configFile := os.Getenv(totemConfigVar)
	if configFile == "" {
		homeDir := os.Getenv("HOME")
		os.Setenv(totemConfigVar, strings.Join(
			[]string{homeDir, pathToConfig, totemConfig}, "/"),
		)
	}
}

func getUserFromConfigs(pathToConfig string) *User {
	setTotemConfigs(pathToConfig)
	user := parseConfigs(os.Getenv(totemConfigVar))
	return &User{user.email, user.password}
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
		panic(err)
	}

	url := os.Getenv("BASE_URL")

	var token string

	for _, arg := range args {

		switch arg {
		case Add:
			add()
		case List:
			list()
		case Auth:
			user := getUserFromConfigs("")
			token = auth.GetToken(&url, &user.email, &user.password)
			fmt.Println(token)
		case Help:
			help()
		case "":
			help()
		default:
			help()
		}
	}
}
