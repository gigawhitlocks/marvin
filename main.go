package main

import (
	"flag"
	bot "github.com/gigawhitlocks/marvin/bot"
)

var (
	password string
	username string
	server   string
)

func init() {
	flag.StringVar(&username, "username", "bot", "bot's username")
	flag.StringVar(&password, "password", "password", "bot's password")
	flag.StringVar(&server, "server", "http://localhost:3000", "server url")
	flag.Parse()
}

func main() {
	s := &bot.Session{}
	s.Logon(username, password, server)
	s.JoinChannel("GENERAL")
	s.Say("GENERAL", "test")
	s.LeaveChannel("GENERAL")
}
