package main

import (
	"dreonbot/pkg/adapters"
	server "dreonbot/pkg/infrastructures"
)

func init() {
	adapters.IocConfigs()
	adapters.IoCLogger()
	adapters.TelebotStart()
}

func main() {
	server.StartEchoServer()
}
