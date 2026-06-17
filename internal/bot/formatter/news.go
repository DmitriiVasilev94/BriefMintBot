package formatter

import (
	"fmt"
	"github.com/DmitriiVasilev94/BriefMintBot/internal/domain"
	"strings"
)

type Formatter interface {
	Format(intro string, emoji string, news []domain.NewsItem) string
}

type HtmlFormatter struct{}

func (f *HtmlFormatter) Format(intro string, emoji string, news []domain.NewsItem) string {
	var builder strings.Builder
	builder.WriteString(intro)
	for index, newsItem := range news {
		line := fmt.Sprintf("%s <a href=\"%s\"><b>%d. %s</b></a>\n<i>%s</i>\n\n",
			emoji,
			newsItem.Link,
			index+1,
			newsItem.Title,
			newsItem.Description,
		)
		builder.WriteString(line)
	}
	return builder.String()
}
