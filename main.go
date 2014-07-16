package main

import (
	clientInterface "github.com/yoed/yoed-client-interface"
	"github.com/yoed/yoed-handler-mailjet/handler"
)

func main() {
	handler := handler.New()
	client := clientInterface.New(handler, &handler.Config.Config)
	client.Run()
}