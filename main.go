package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

var (
	password string
	username string
	server   string

	userId    string
	authToken string
)

type logonData struct {
	AuthToken string
	UserId    string
}

type logonResponse struct {
	Status string
	Data   logonData
}

func init() {
	flag.StringVar(&username, "username", "bot", "bot's username")
	flag.StringVar(&password, "password", "password", "bot's password")
	flag.StringVar(&server, "server", "https://localhost:3000", "server url")
	flag.Parse()
}

func logon() (string, string) {
	b := new(bytes.Buffer)

	b.WriteString("user=")
	b.WriteString(username)
	b.WriteString("&password=")
	b.WriteString(password)

	resp, _ := http.Post(server+"/api/login",
		"application/x-www-form-urlencoded",
		b)

	var l logonResponse
	err := json.NewDecoder(resp.Body).Decode(&l)
	if err != nil {
		panic("Probably not the right username/password/server")
	}

	return l.Data.UserId, l.Data.AuthToken
}

func joinChannel(channelName string) {

}

func main() {
	userId, authToken := logon()
	fmt.Printf("%s %s", userId, authToken)
	joinChannel("general")
}
