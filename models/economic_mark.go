package models

import (
	"time"

	"github.com/chaso-pa/ecomimic-cron/middleware"
	"gorm.io/gorm/clause"
)

// EconomicMark represents the economic_marks table in the database.
type EconomicMark struct {
	ID          string    `json:"id" db:"id"`
	PublishedAt time.Time `json:"publishedAt" db:"published_at"`
	Country     string    `json:"country" db:"country"`
	Title       string    `json:"title" db:"title"`
	Importance  int       `json:"importance" db:"importance"`
	Estimate    string    `json:"estimate" db:"estimate"`
	Result      string    `json:"result" db:"result"`
	PastResult  string    `json:"pastresult" db:"past_result"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

func EconomicMarkUpsert(marks []*EconomicMark) {
	middleware.GetDb().Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{`published_at`, `country`, `title`, `importance`, `estimate`, `result`, `past_result`, `created_at`, `updated_at`}),
	}).Create(marks)
}
