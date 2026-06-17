package bot

import (
	"context"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot/formatter"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/provider/news"
	tgBot "github.com/go-telegram/bot"
	"log"
)

type BriefMintBot struct {
	Client            *tgBot.Bot
	worldNewsProvider news.NewsProvider
	techNewsProvider  news.NewsProvider
	newsFormatter     formatter.Formatter
}

func New(tgToken string, worldNewsProvider news.NewsProvider, techNewsProvider news.NewsProvider, newsFormatter formatter.Formatter) *BriefMintBot {
	bb := &BriefMintBot{
		worldNewsProvider: worldNewsProvider,
		techNewsProvider:  techNewsProvider,
		newsFormatter:     newsFormatter,
	}

	opts := []tgBot.Option{
		tgBot.WithDefaultHandler(bb.mainHandler),
	}

	client, err := tgBot.New(tgToken, opts...)
	if err != nil {
		log.Println("error during bot initialization: " + err.Error())
		panic(err)
	}
	return &BriefMintBot{
		Client:            client,
		worldNewsProvider: worldNewsProvider,
	}
}

func (b *BriefMintBot) Start(ctx context.Context) {
	b.Client.Start(ctx)
}
