package bot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)


func (bb *BriefMintBot) mainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		log.Println("update.Message is nil, cannot process it")
		return
	}
		
	if update.Message.Text == "" {
		log.Printf("User: %v sent unsupported content", update.Message.From.Username)
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Something goes wrong. The message content is not supported",
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}	
		return
	}

	if update.Message.Text == "/brief" {
		log.Printf("brief is requested by %v", update.Message.From.Username)
		titles, err := bb.newsProvider.FetchRecentNews(ctx)
		if err != nil {
			log.Printf("News provider returns error: %v", err.Error())
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Can't prepare brief, try again later.",
			})
			if err != nil {
				log.Printf("Failed to send message: %v", err)
			}
			return
		}
		var builder strings.Builder 
		for index, title := range titles {
			builder.WriteString(fmt.Sprintf("#%v. %s\n", index+1, title))
			if index==15 {
				break
			}
		}
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   builder.String(),
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
		return
		
	}
	
	log.Printf("Got message: \"%v\" from: %v",update.Message.Text, update.Message.From.Username)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}