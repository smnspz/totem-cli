package auth

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/term"
)

func getPassword() string {
	// https://groups.google.com/g/golang-codereviews/c/CM9H5ya03lE?pli=1
	tty, err := os.Open("/dev/tty")
	if err != nil {
		panic(err)
	}
	defer tty.Close()
	fmt.Print("Type your password ")
	pwd, err := term.ReadPassword(int(tty.Fd()))
	if err != nil {
		panic(err)
	}
	return string(pwd)
}

func getUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Trim(username, "\n")
}

func GetToken(baseUrl string) string {
	// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
	username := getUsername()
	password := getPassword()

	body, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	if err != nil {
		log.Fatalln(err)
	}

	url := baseUrl + "/jwt/login"

	resp, httpErr := http.Post(url, "application/json", bytes.NewBuffer(body))
	if httpErr != nil {
		log.Fatalln(httpErr)
	}

	defer resp.Body.Close()

	body, ioErr := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(ioErr)
	}

	// log.Println(string(body))
	return string(body)
}
