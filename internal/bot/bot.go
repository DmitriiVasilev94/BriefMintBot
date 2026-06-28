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
	newsProvider news.NewsProvider
	newsFormatter     formatter.Formatter
}

func New(tgToken string, newsProvider news.NewsProvider, newsFormatter formatter.Formatter) *BriefMintBot {
	bb := &BriefMintBot{
		newsProvider: newsProvider,
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
		newsProvider: newsProvider,
	}
}

func (b *BriefMintBot) Start(ctx context.Context) {
	b.Client.Start(ctx)
}
