package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	return &Config{
		TelegramBotToken: telegramToken,
	}
	
}