package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/config"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/provider/news"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.LoadConfig()

	client := bot.New(cfg.TelegramBotToken, &news.BBCProvider{URL: "https://news.google.com/rss/search?q=Global%20business&hl=en-US&gl=US&ceid=US%3Aen"})

	log.Println("Bot started successfully")
	client.Start(ctx)
}
