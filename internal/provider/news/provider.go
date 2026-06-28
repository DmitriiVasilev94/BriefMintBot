package news

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/DmitriiVasilev94/BriefMintBot/internal/domain"
	"github.com/mmcdole/gofeed"
)

const maxNewsFromSource = 5
const cacheTTLInSeconds = 30

type NewsProvider interface {
	FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error)
}

type RSSProvider struct {
	URLs []string
	cache sync.Map
	mutex sync.Mutex
}

type CachedNews struct {
	Items []domain.NewsItem
	LastUsed time.Time
}

func (p *RSSProvider) FetchRecentNews(ctx context.Context) ([]domain.NewsItem, error) {
	var news []domain.NewsItem
	
	for _, url := range p.URLs {

		items, ok := p.getCachedItems(url)
		if ok {
			news = append(news, items...)
			continue
		}
		
		p.mutex.Lock()
		items, ok = p.getCachedItems(url)
		if ok {
			news = append(news, items...)
			continue
		}
		
		items, err := p.fetchNewsFromUrl(ctx, url, maxNewsFromSource)
		if err != nil {
			log.Printf("Can't fetch news from URL %s: %v", url, err.Error())
			return nil, err
		}
		p.cache.Store(url, &CachedNews{
			Items: items,
			LastUsed: time.Now(),
		})
		p.mutex.Unlock()
		log.Printf("Logs for %s updated in cache.", url)

		news = append(news, items...)
	}
	return news, nil
}

func (p *RSSProvider) getCachedItems(url string) ([]domain.NewsItem, bool) {
	cached, ok := p.cache.Load(url)
	if ok {
		cachedNews := cached.(*CachedNews)
		if time.Since(cachedNews.LastUsed) <=cacheTTLInSeconds * time.Second {
			log.Printf("Logs for %s loaded from cache.", url)
			return  cachedNews.Items, true
		}
		p.cache.Delete(url)
		log.Printf("Logs for %s expired, clear the cache.", url)

	}
	return nil, false
}

func (p *RSSProvider) fetchNewsFromUrl(ctx context.Context, url string, maxNews int) ([]domain.NewsItem, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Printf("Can't parse news: %v", err.Error())
		return nil, err
	}
	var items []domain.NewsItem
	for index, item := range feed.Items {
		items = append(items, domain.NewsItem{Title: item.Title, Link: item.Link, Description: item.Description})
		if index+1 >=maxNewsFromSource {
			break
		}
	}
	return items,nil
	
}
