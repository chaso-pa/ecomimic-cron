package models

import (
	"github.com/chaso-pa/ecomimic-cron/middleware"
	"gorm.io/gorm/clause"
	"time"
)

type Article struct {
	ID          string    `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ContentHash string    `json:"content_hash"`
	Published   time.Time `json:"published"`
	CrawledAt   time.Time `json:"crawled_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ArtcileUpsert(articles []*Article) {
	middleware.GetDb().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{`url`, `title`, `content`, `content_hash`, `published`, `crawled_at`, `created_at`, `updated_at`}),
	}).Create(articles)
}
