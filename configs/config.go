package configs

import (
	"os"

	"github.com/golobby/dotenv"
)

type AppConfig struct {
	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}
	Mongo struct {
		Url string `env:"MONGO_URL"`
		Db  string `env:"MONGO_DB"`
	}

	GenAI struct {
		GeminiKey string `env:"GENAI_GEMINI_KEY"`
	}

	Telegram struct {
		BotToken string `env:"TELE_BOT_TOKEN"`
	}
}

func NewAppConfig(envDir string) (*AppConfig, error) {
	appConfig := &AppConfig{}
	file, err := os.Open(envDir)
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(appConfig)
	if err != nil {
		return nil, err
	}

	return appConfig, nil
}
