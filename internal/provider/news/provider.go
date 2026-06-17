package news

import (
	"context"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/domain"
	"github.com/mmcdole/gofeed"
	"log"
)

type NewsProvider interface {
	FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error)
}

type RSSProvider struct {
	URL string
}

func (p *RSSProvider) FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(p.URL)
	if err != nil {
		log.Printf("Can't parse news: %v", err.Error())
		return nil, err
	}
	var news []domain.NewsItem
	for index, item := range feed.Items {
		news = append(news, domain.NewsItem{Title: item.Title, Link: item.Link, Description: item.Description})
		if index == 4 {
			break //no more than 5 news
		}
	}
	return news, nil
}
