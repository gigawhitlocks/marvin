package bot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type logonData struct {
	UserId    string
	AuthToken string
}

type result struct {
	Status string
}

type logonResponse struct {
	Status string
	Data   logonData
}

type Session struct {
	logonData
	server string
}

func (s *Session) Logon(username, password, server string) {
	s.server = server

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

	s.UserId = l.Data.UserId
	s.AuthToken = l.Data.AuthToken
}

func (s *Session) authRequest(url string) (*http.Request, *http.Client) {
	client := &http.Client{}
	b := new(bytes.Buffer)
	req, _ := http.NewRequest("POST", url, b)

	req.Header.Add("X-Auth-Token", s.AuthToken)
	req.Header.Add("X-User-Id", s.UserId)

	return req, client
}

func (s *Session) JoinChannel(channelName string) {
	req, client := s.authRequest(s.server + "/api/rooms/" + channelName + "/join")
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
}

func (s *Session) Say(channelName, message string) {
	client := &http.Client{}
	b := new(bytes.Buffer)
	b.WriteString("{\"msg\": \"" + message + "\"}")
	req, _ := http.NewRequest("POST",
		s.server+"/api/rooms/"+channelName+"/send", b)

	req.Header.Add("X-Auth-Token", s.AuthToken)
	req.Header.Add("X-User-Id", s.UserId)
	req.Header.Add("Content-Type", "application/json")

	client.Do(req)
}

func (s *Session) LeaveChannel(channelName string) {
	req, client := s.authRequest(s.server + "/api/rooms/" + channelName + "/leave")
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
}
