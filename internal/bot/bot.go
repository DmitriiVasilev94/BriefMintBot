package bot

import (
	 "log"
	 "context"
	 tgBot "github.com/go-telegram/bot"
)

type Bot struct {
	Client *tgBot.Bot

}

func New(tgToken string) *Bot {

	opts := []tgBot.Option{
		tgBot.WithDefaultHandler(mainHandler),
	}
	
	client, err := tgBot.New(tgToken, opts...)
	if err != nil {
		log.Println("error during bot initialization: " + err.Error())
		panic(err)
	}
	return &Bot{
		Client: client,
	}
}

func (b *Bot) Start(ctx context.Context){
	b.Client.Start(ctx)
}