package adapters

import (
	"dreonbot/configs"
	"dreonbot/libs/telebot"
	"dreonbot/shared/interfaces"
	"log"

	"github.com/golobby/container/v3"
)

func TelebotStart() {
	var (
		logger    interfaces.ILogger
		appConfig *configs.AppConfig
		err       error
	)

	err = container.Resolve(&logger)
	if err != nil {
		log.Fatal(err)
	}
	err = container.Resolve(&appConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Start your telebot here
	myTelebot := telebot.NewTelebot(appConfig, logger)
	go myTelebot.Start()
}
