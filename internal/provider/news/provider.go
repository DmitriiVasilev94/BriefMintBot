package news

import (
	"context"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/domain"
	"github.com/mmcdole/gofeed"
	"log"
)

const maxNewsFromSource = 5

type NewsProvider interface {
	FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error)
}

type RSSProvider struct {
	URLs []string
}

func (p *RSSProvider) FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error) {
	fp := gofeed.NewParser()
	var news []domain.NewsItem
	for _, url := range p.URLs {
		feed, err := fp.ParseURL(url)
		if err != nil {
			log.Printf("Can't parse news: %v", err.Error())
			return nil, err
		}
		for index, item := range feed.Items {
			news = append(news, domain.NewsItem{Title: item.Title, Link: item.Link, Description: item.Description})
			if index+1 >=maxNewsFromSource {
				break
			}
		}
	}
	return news, nil
}
