package bot

import (
	"context"
	"log"

	"github.com/DmitriiVasilev94/BriefMintBot/internal/bot/formatter"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/provider/news"
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
		sendNewsMessage(ctx, b, "Here is latest 5 world news from BBC:\n", "🌍", bb.worldNewsProvider, bb.newsFormatter, update.Message.Chat.ID)
		sendNewsMessage(ctx, b, "And take  latest 5 world news from TechCrunch.com:\n", "📡", bb.techNewsProvider, bb.newsFormatter, update.Message.Chat.ID)

		return

	}
	//default behaivor
	log.Printf("Got message: \"%v\" from: %v", update.Message.Text, update.Message.From.Username)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello! There is only one supported command /brief",
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

func sendNewsMessage(ctx context.Context,
	b *bot.Bot,
	intro string,
	emoji string,
	provider news.NewsProvider,
	formatter formatter.Formatter, chatId int64) {
	news, err := provider.FetchRecentNews(ctx)
	if err != nil {
		log.Printf("News provider returns error: %v", err.Error())
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Can't prepare brief, try again later.",
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
		return
	}
	var newsMsg = formatter.Format(intro, emoji, news)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatId,
		Text:      newsMsg,
		ParseMode: "HTML",
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
