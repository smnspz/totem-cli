package auth

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func GetToken(baseUrl *string, username *string, password *string) string {
	// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
	body, err := json.Marshal(map[string]string{
		"username": *username,
		"password": *password,
	})

	if err != nil {
		fmt.Println("Error while parsing user json")
		panic(err)
	}

	url := *baseUrl + "/jwt/login"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error while sending request")
		panic(err)
	}

	defer resp.Body.Close()

	body, ioErr := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response")
		panic(ioErr)
	}

	var retVal map[string]interface{}

	if err := json.Unmarshal(body, &retVal); err != nil {
		fmt.Println("Error while parsing response json")
		panic(err)
	}

	return retVal["token"].(string)
}
