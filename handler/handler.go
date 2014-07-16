package handler

import (
	httpInterface "github.com/yoed/yoed-http-interface"
	"net/http"
	"net/url"
	"log"
	"strings"
)

type Handler struct {
	Config *Config
}

type Config struct {
	httpInterface.Config
	ApiKey string
	ApiSecret string
	FromEmail string
	ToEmail string
	Subject string
	Text string
}

func (c *Handler) Handle(username string) {
	client := &http.Client{}

	data := url.Values {
		"from":{c.Config.FromEmail},
		"to": {c.Config.ToEmail},
		"subject": {strings.Replace(c.Config.Subject, "%username%", username, -1)},
		"text": {strings.Replace(c.Config.Text, "%username%", username, -1)},
	}

	log.Printf("Yo'ed by %s: sending an email to %s", username, c.Config.ToEmail)

	req, _ := http.NewRequest("POST", "https://api.mailjet.com/v3/send/message", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.Config.ApiKey, c.Config.ApiSecret)
	client.Do(req)
}

func New() *Handler {

	c := &Handler{}

	if err := httpInterface.LoadConfig("./config.json", &c.Config); err != nil {
		log.Fatalf("failed loading config: %s", err)
	}

	return c
}