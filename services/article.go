package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/chaso-pa/ecomimic-cron/models"
	"github.com/gocolly/colly"
	"github.com/lucsky/cuid"
)

// fetch all crawlable article domains and fetch articles
func CrawlAllArticles() error {
	articleDomains := models.CrawlableArticleDomains()
	for _, ad := range articleDomains {
		err := fetchArticles(ad)
		if err != nil {
			return err
		}
	}

	return nil
}

func fetchArticles(ad *models.ArticleDomain) error {
	var err error
	c := setupColly(ad.Domain)

	c.OnHTML(ad.ContainerSelector, func(e *colly.HTMLElement) {
		e.ForEach(ad.ArticleLinkSelector, func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Attr("href"))
			time.Sleep(10 * time.Second)
			err = saveArticle(el.Attr("href"), ad)
			if err != nil {
				return
			}
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		err = errors.New("Race Crawl Error")
	})

	// Start the collector
	err = c.Visit(ad.URL)
	return err
}

func saveArticle(url string, ad *models.ArticleDomain) error {
	var err error
	c := setupColly(ad.Domain)

	var articles []*models.Article
	c.OnHTML(ad.ArticleContainerSelector, func(e *colly.HTMLElement) {
		title := e.ChildText(ad.TitleSelector)
		content := e.ChildText(ad.ContentSelector)
		articleURL := e.Request.URL.String()
		publishedAt, err := parseTime(e.ChildAttr(ad.PublishedAtSelector, "datetime"))
		if err != nil {
			publishedAt, err = parseTime(e.ChildText(ad.PublishedAtSelector))
		}
		if err != nil {
			publishedAt = time.Now()
		}

		article := models.Article{
			ID:          cuid.New(),
			URL:         articleURL,
			Title:       title,
			Content:     content,
			ContentHash: HashString(title),
			Published:   publishedAt,
			CrawledAt:   time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Save article to database
		articles = append(articles, &article)
		models.ArtcileUpsert(articles)
	})

	c.OnError(func(r *colly.Response, e error) {
		err = errors.New("Race Crawl Error")
	})

	// Start the collector
	err = c.Visit(ad.ArticleUrlBase + url)

	return err
}

// can parse any time string format
func parseTime(timeStr string) (time.Time, error) {
	var err error
	var t time.Time
	timeFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, timeFormat := range timeFormats {
		t, err = time.Parse(timeFormat, timeStr)
		if err == nil {
			return t, nil
		}
	}

	return t, err
}
