package bot

import 	(
	"log"
	"context"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/bot"
)


func mainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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
	log.Printf("Got message: \"%v\" from: %v",update.Message.Text, update.Message.From.Username)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}