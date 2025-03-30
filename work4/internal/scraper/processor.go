package scraper

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type Processor struct {
	Timeout    time.Duration
	RetryCount int
}

type Result struct {
	Date        time.Time
	URL         string
	Title       string
	Description string
	StatusCode  int
}

type pageInfo struct {
	Title       string
	Description string
}

func NewProcessor(conf *Config) *Processor {
	return &Processor{Timeout: conf.Timeout, RetryCount: conf.RetryCount}
}

func (p *Processor) Do(in <-chan string, out chan<- *Result) {
	for url := range in {
		res, err := p.get(url)
		if err != nil {
			slog.Error("processing error", slog.String("url", url), slog.String("err", err.Error()))
			continue
		}

		if res.StatusCode >= 400 { //nolint:mnd
			slog.Error("processing error", slog.String("url", url), slog.String("err", res.Status), slog.Int("status_code", res.StatusCode))
			continue
		}

		info, err := p.parse(res)
		if err != nil {
			slog.Error("processing error", slog.String("url", url), slog.String("err", err.Error()))
			continue
		}

		out <- &Result{
			Date:        time.Now(),
			URL:         url,
			StatusCode:  res.StatusCode,
			Title:       info.Title,
			Description: info.Description,
		}
	}
}

func (p *Processor) get(url string) (*http.Response, error) {
	client := http.Client{Timeout: p.Timeout}

	req, err := http.NewRequest(http.MethodGet, url, nil) //nolint:noctx
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	var res *http.Response

	for range p.RetryCount {
		if res, err = client.Do(req); err != nil {
			return nil, fmt.Errorf("response error: %w", err)
		}

		if res.StatusCode == http.StatusServiceUnavailable {
			slog.Warn("page is temporarily unavailable, retrying")
			continue
		}

		return res, nil
	}

	return res, nil
}

func (p *Processor) parse(res *http.Response) (*pageInfo, error) {
	ct := res.Header.Get("Content-Type")
	defer res.Body.Close()

	reader, err := charset.NewReader(res.Body, ct)
	if err != nil {
		return nil, fmt.Errorf("charset detection error: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("document read error: %w", err)
	}

	info := pageInfo{}

	doc.Find("head").Each(func(_ int, s *goquery.Selection) {
		info.Title = s.Find("title").Text()

		value, exist := s.Find("meta[name='description']").Attr("content")
		if exist {
			info.Description = value
		}
	})

	return &info, nil
}

func (rs *Result) Fields() []string {
	return []string{
		rs.Date.Format(time.RFC3339),
		rs.URL,
		strconv.Itoa(rs.StatusCode),
		rs.Title,
		rs.Description,
	}
}
