package models

import (
	"time"

	"github.com/chaso-pa/ecomimic-cron/middleware"
	"gorm.io/gorm"
)

type ArticleDomain struct {
	ID                       string    `json:"id"`
	Domain                   string    `json:"domain"`
	URL                      string    `json:"url"`
	ArticleUrlBase           string    `json:"article_url_base"`
	CrawlStatus              bool      `json:"crawl_status"`
	ContainerSelector        string    `json:"container_selector"`
	ArticleLinkSelector      string    `json:"article_link_selector"`
	ArticleContainerSelector string    `json:"article_container_selector"`
	TitleSelector            string    `json:"title_selector"`
	ContentSelector          string    `json:"content_selector"`
	PublishedAtSelector      string    `json:"published_at_selector"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func AllArticleDomains() []*ArticleDomain {
	var articleDomains []*ArticleDomain
	middleware.GetDb().Find(&articleDomains)
	return articleDomains
}

func CrawlableArticleDomains() []*ArticleDomain {
	var articleDomains []*ArticleDomain
	middleware.GetDb().Scopes(crawlable()).Find(&articleDomains)
	return articleDomains
}

func crawlable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("crawl_status = 1")
	}
}
