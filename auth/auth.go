package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetToken(username string, password string, baseUrl string) {
	// https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
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

	log.Println(string(body))
}
