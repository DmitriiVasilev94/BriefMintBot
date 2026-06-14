package main

import (
	"context"
	"os"
	"os/signal"
	"log"

	"github.com/DmitriiVasilev94/BriefMintBot/internal/config"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.LoadConfig()

	client := bot.New(cfg.TelegramBotToken)
	
	log.Println("Bot started successfully")
	client.Start(ctx)
}