package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot/formatter"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/config"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/provider/news"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.LoadConfig()

	client := bot.New(cfg.TelegramBotToken,
		&news.RSSProvider{URL: "https://feeds.bbci.co.uk/news/world/rss.xml"},
		&news.RSSProvider{URL: "https://techcrunch.com/feed/"},
		&formatter.HtmlFormatter{})

	log.Println("Bot started successfully")
	client.Start(ctx)
}
