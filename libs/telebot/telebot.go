package telebot

import (
	"dreonbot/configs"
	"dreonbot/libs/gemini"
	"dreonbot/shared/interfaces"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Telebot struct {
	configs *configs.AppConfig
	logger  interfaces.ILogger
	bot     *tele.Bot
	gemimi  *gemini.GenAIGemini
}

func NewTelebot(configs *configs.AppConfig, logger interfaces.ILogger) *Telebot {
	settings := tele.Settings{
		Token: configs.Telegram.BotToken,
		Poller: &tele.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	bot, err := tele.NewBot(settings)
	if err != nil {
		logger.Error(err.Error())
	}

	gemimi, err := gemini.NewGenAIGemini(configs.GenAI.GeminiKey)
	if err != nil {
		logger.Error(err.Error())
	}

	return &Telebot{
		configs: configs,
		logger:  logger,
		bot:     bot,
		gemimi:  gemimi,
	}
}

func (t *Telebot) Start() {
	t.bot.Handle(tele.OnText, func(c tele.Context) error {

		var (
			user = c.Sender()
			text = c.Text()
		)

		t.logger.InfoF("User: %s, Text: %s", user.Username, text)

		// Generate content
		resp, err := t.gemimi.GenerateContent(text)
		if err != nil {
			t.logger.Error(err.Error())
			return c.Send("Sorry, I can't generate content right now.")
		}

		// Send the generated content
		return c.Send(resp)
	})

	t.logger.Info("Telebot is running...")

	t.bot.Start()
}
