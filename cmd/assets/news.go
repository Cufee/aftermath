package main

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

var httpClient = http.Client{
	Timeout: time.Second * 15,
}

type Image struct {
	image.Image
	PublishDate time.Time
	Name        string
	URL         string
}

type NewsItem struct {
	Name   string `json:"name"`
	Images struct {
		Medium string `json:"medium"`
	} `json:"images"`
	PublishDate       string    `json:"publication_start_at"`
	PublishDateParsed time.Time `json:"-"`
}

func ScrapeNewsImages() error {
	endpoint := os.Getenv("NEWS_ENDPOINT")
	if endpoint == "" {
		return errors.New("NEWS_ENDPOINT not set")
	}

	link, err := url.Parse(endpoint)
	if err != nil {
		return errors.Wrap(err, "invalid endpoint url")
	}

	images, err := scrapeSinceDate(link, time.Now().Add(time.Hour*24*-1), 0)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	validImagesCh := make(chan Image, len(images))
	for _, img := range images {
		wg.Add(1)
		go func(img NewsItem) {
			defer wg.Done()

			res, err := httpClient.Get(img.Images.Medium)
			if err != nil {
				log.Err(err).Msg("failed to get image")
				return
			}
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			if err != nil {
				log.Err(err).Msg("failed to read image response body")
				return
			}

			decoded, err := imaging.Decode(bytes.NewReader(data))
			if err != nil {
				log.Err(err).Msg("failed to decode image")
				return
			}
			validImagesCh <- Image{decoded, img.PublishDateParsed, img.Name, img.Images.Medium}

		}(img)
	}
	wg.Wait()
	close(validImagesCh)

	for img := range validImagesCh {
		println("image", img.Name)
	}

	return nil
}

func scrapeSinceDate(endpoint *url.URL, sinceDate time.Time, page int) ([]NewsItem, error) {
	query := url.Values{}
	query.Set("page", strconv.Itoa(page))
	query.Set("page_size", "10")
	endpoint.RawQuery = query.Encode()

	var data struct {
		Count   int        `json:"count"`
		Results []NewsItem `json:"results"`
	}

	res, err := httpClient.Get(endpoint.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var validImages []NewsItem
	var oldestDate = time.Now()
	for _, img := range data.Results {
		img.PublishDateParsed, err = time.Parse("2006-01-02T15:04:05", img.PublishDate)
		if err != nil {
			log.Err(err).Msg("failed to parse date")
			continue
		}

		if img.PublishDateParsed.Before(sinceDate) {
			oldestDate = img.PublishDateParsed
			break
		}
		if strings.Contains(strings.ToLower(img.Name), "video:") {
			continue
		}

		validImages = append(validImages, img)
		if img.PublishDateParsed.Before(oldestDate) {
			oldestDate = img.PublishDateParsed
		}
	}

	if oldestDate.Before(sinceDate) {
		return validImages, nil
	}

	nextPage, err := scrapeSinceDate(endpoint, sinceDate, page+1)
	return append(validImages, nextPage...), err
}
