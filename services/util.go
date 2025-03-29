package services

import (
	"archive/zip"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
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

func DeleteDailyContents() error {
	paths, err := filepath.Glob("ai_gens/202*")
	if err != nil {
		return err
	}

	for _, path := range paths {
		os.RemoveAll(path)
	}

	return nil
}

func ZipDir(dirPath string) error {
	zipFile, err := os.Create(dirPath + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	files, err := filepath.Glob(dirPath + "/*")
	if err != nil {
		return err
	}

	for _, filename := range files {
		if err := AddFileToZip(filename, zipWriter); err != nil {
			return err
		}
	}

	return nil
}

func AddFileToZip(filename string, zipWriter *zip.Writer) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
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
