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
	Server struct {
		Host string `env:"HTTP_HOST"`
		Port string `env:"HTTP_PORT"`
	}

	Auth struct {
		ignoreMethods []string `env:"AUTH_IGNORE_METHODS"`
		JWT           struct {
			SecretKey        string `env:"JWT_SECRET_KEY"`
			RefreshSecretKey string `env:"JWT_REFRESH_SECRET_KEY"`
			Issuer           string `env:"JWT_ISSUER"`
			ExpiredTime      int64  `env:"JWT_EXPIRED_TIME"`
			RefreshExpired   int64  `env:"JWT_REFRESH_EXPIRED_TIME"`
		}
		IgnoreMethods map[string]bool
	}

	GenAI struct {
		GeminiKey string `env:"GENAI_GEMINI_KEY"`
	}

	Telegram struct {
		BotName  string `env:"TELE_BOT_NAME"`
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
