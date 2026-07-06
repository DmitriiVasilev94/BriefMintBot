package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
	LLMStudioUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	llmStudioUrl := os.Getenv("LM_STUDIO_URL")
	return &Config{
		TelegramBotToken: telegramToken,
		LLMStudioUrl: llmStudioUrl,
	}

}
