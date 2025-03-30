package services

import (
	"crypto/sha256"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"time"
)

const UA string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"

func setupColly(ad string) colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(UA),
	)

	c.SetRequestTimeout(15 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL.String())
	})
	c.Limit(&colly.LimitRule{
		DomainGlob:  (ad),
		Delay:       1 * time.Second,
		RandomDelay: 3 * time.Second,
	})
	return *c
}

func GetQueryParams(rawUrl string) (map[string]string, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("URL parsing error: %v", err)
	}
	result := make(map[string]string)
	for k, v := range parsedURL.Query() {
		result[k] = v[0]
	}
	return result, nil
}

func RemoveLastChar(input string) string {
	// 文字列の長さが1以上の場合、最後の文字を削除
	if len(input) > 0 {
		return input[:len(input)-1]
	}
	return input
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
