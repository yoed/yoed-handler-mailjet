package main

import (
	clientInterface "github.com/yoed/yoed-client-interface"
	"net/http"
	"net/url"
	"log"
	"fmt"
	"encoding/json"
	"strings"
)

type MailjetYoedClient struct {
	clientInterface.BaseYoedClient
	config *MailjetYoedClientConfig
}

type MailjetYoedClientConfig struct {
	ApiKey string
	ApiSecret string
	FromEmail string
	ToEmail string
	Subject string
	Text string
}

func (c *MailjetYoedClient) loadConfig(configPath string) (*MailjetYoedClientConfig, error) {
	configJson, err := clientInterface.ReadConfig(configPath)

	if err != nil {
		return nil, err
	}

	config := &MailjetYoedClientConfig{}

	if err := json.Unmarshal(configJson, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *MailjetYoedClient) Handle(username string) {
	client := &http.Client{}

	data := url.Values {
		"from":{c.config.FromEmail},
		"to": {c.config.ToEmail},
		"subject": {strings.Replace(c.config.Subject, "%username%", username, -1)},
		"text": {strings.Replace(c.config.Text, "%username%", username, -1)},
	}

	log.Printf("Yo'ed by %s: sending an email to %s", username, c.config.ToEmail)

	req, _ := http.NewRequest("POST", "https://api.mailjet.com/v3/send/message", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.config.ApiKey, c.config.ApiSecret)
	client.Do(req)
}


func NewMailjetYoedClient() (*MailjetYoedClient, error) {
	c := &MailjetYoedClient{}
	config, err := c.loadConfig("./config.json")

	if err != nil {
		panic(fmt.Sprintf("failed loading config: %s", err))
	}

	c.config = config
	baseClient, err := clientInterface.NewBaseYoedClient()

	if err != nil {
		return nil, err
	}
	c.BaseYoedClient = *baseClient

	return c, nil
}

func main() {
	c, _ := NewMailjetYoedClient()

	clientInterface.Run(c)
}