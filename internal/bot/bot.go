package bot

import (
	 "log"
	 "context"
	 tgBot "github.com/go-telegram/bot"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/provider/news"
)

type BriefMintBot struct {
	Client *tgBot.Bot
	newsProvider news.NewsProvider

}

func New(tgToken string, newsProvider news.NewsProvider) *BriefMintBot {
	bb := &BriefMintBot{
        newsProvider: newsProvider,
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
		Client: client,
		newsProvider: newsProvider,
	}
}

func (b *BriefMintBot) Start(ctx context.Context){
	b.Client.Start(ctx)
}