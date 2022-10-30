package ting

import (
	"github.com/nhk-news-web-easy/nhk-news-fetcher-go"
	"log"
	"time"
)

const programId = 1

func RunTask() error {
	log.Println("Start to fetch news")

	now := time.Now().UTC()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	newsList, err := nhk_fetcher.FetchNews(startDate, endDate)

	if err != nil {
		log.Printf("Fetch news error, %v\n", err)

		return err
	}

	if newsList == nil || len(newsList) == 0 {
		log.Println("No news found")

		return nil
	} else {
		log.Printf("Found %v news\n", len(newsList))
	}

	dbConfig, err := ParseDbConfig()

	if err != nil {
		return err
	}

	err = InitDb(dbConfig)

	if err != nil {
		return err
	}

	defer CloseDb()

	for _, news := range newsList {
		ting := Ting{
			ProgramId:   programId,
			Title:       news.Title,
			Description: news.Outline,
			AudioUrl:    news.M3u8Url,
			Content:     news.BodyWithoutHtml,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		err = saveTing(ting)

		if err != nil {
			return err
		}
	}

	return nil
}
