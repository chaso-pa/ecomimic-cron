package services

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chaso-pa/ecomimic-cron/models"
	"github.com/gocolly/colly"
	"github.com/lucsky/cuid"
)

func CrawlEconomicCalendar() error {
	var err error
	var marks []*models.EconomicMark
	targetUrl := "https://www.gaikaex.com/gaikaex/mark/calendar/"
	c := setupColly("gaikaex.com")

	c.OnHTML("#mainContents", func(e *colly.HTMLElement) {
		var date string
		e.ForEach(".tableA01.pcOnly tbody tr", func(_ int, el *colly.HTMLElement) {
			cls := el.Attr("class")
			if cls == "holiday" {
				return
			}

			firstTd := el.ChildText("td:first-child")
			if ValidateDateTime(firstTd) {
				date = el.ChildText("td:nth-child(1)")
				id := el.Attr("id")
				timeStr := el.ChildText("td:nth-child(2)")
				publishedAt, _ := parseDateTime(date, timeStr, id)
				country := el.ChildAttr("td:nth-child(3) img", "alt")
				title := el.ChildText("td:nth-child(4)")
				importance, _ := parseImportance(el.ChildText("td:nth-child(5)"))
				estimate := el.ChildText("td:nth-child(6)")
				result := el.ChildText("td:nth-child(7)")
				pastresult := el.ChildText("td:nth-child(8)")

				mark := models.EconomicMark{
					ID:          cuid.New(),
					PublishedAt: publishedAt,
					Country:     country,
					Title:       title,
					Importance:  int(importance),
					Estimate:    estimate,
					Result:      result,
					PastResult:  pastresult,
				}
				marks = append(marks, &mark)

			} else {
				id := el.Attr("id")
				timeStr := el.ChildText("td:nth-child(1)")
				publishedAt, _ := parseDateTime(date, timeStr, id)
				country := el.ChildText("td:nth-child(2)")
				title := el.ChildText("td:nth-child(3)")
				importance, _ := parseImportance(el.ChildText("td:nth-child(4)"))
				estimate := el.ChildText("td:nth-child(5)")
				result := el.ChildText("td:nth-child(6)")
				pastresult := el.ChildText("td:nth-child(7)")

				mark := models.EconomicMark{
					ID:          cuid.New(),
					PublishedAt: publishedAt,
					Country:     country,
					Title:       title,
					Importance:  int(importance),
					Estimate:    estimate,
					Result:      result,
					PastResult:  pastresult,
				}
				marks = append(marks, &mark)
			}
		})
		models.EconomicMarkUpsert(marks)
	})

	c.OnError(func(r *colly.Response, e error) {
		err = errors.New("crawl error")
	})

	// Start the collector
	err = c.Visit(targetUrl)
	return err
}

// ParseDateTime parses date, time, and id to return time.Time with JST timezone
func parseDateTime(date, timeStr, id string) (time.Time, error) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	// Extract year from id (yyyydd+increment+_emg format)
	if len(id) < 6 {
		return time.Time{}, fmt.Errorf("invalid id format")
	}

	yearStr := id[:4]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year in id: %v", err)
	}

	// Parse date (m/d(w) format)
	datePattern := `^(\d{1,2})/(\d{1,2})\([月火水木金土日]\)$`
	re := regexp.MustCompile(datePattern)
	matches := re.FindStringSubmatch(date)
	if len(matches) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	month, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %v", err)
	}

	day, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %v", err)
	}

	// Parse time (hh:mm format)
	timeParts := strings.Split(timeStr, ":")
	if len(timeParts) != 2 {
		return time.Time{}, fmt.Errorf("invalid time format")
	}

	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid hour: %v", err)
	}

	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid minute: %v", err)
	}

	// Create time.Time with JST timezone
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, jst), nil
}

func parseImportance(importance string) (int32, error) {
	switch importance {
	case "★★★":
		return 3, nil
	case "★★":
		return 2, nil
	case "★":
		return 1, nil
	default:
		return 1, fmt.Errorf("invalid importance: %v", importance)
	}
}
