package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Generator   string `xml:"generator"`
		Language    string `xml:"language"`
		Image       struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			URL   string `xml:"url"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Item []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func urlToFeed(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}
	rssFeed := RssFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RssFeed{}, err
	}

	return rssFeed, nil
}
