package main

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UrlShortenerService
}

func (mw loggingMiddleware) Wrap(svc UrlShortenerService) UrlShortenerService {
	logger := log.NewLogfmtLogger(os.Stderr)

	return loggingMiddleware{next: svc, logger: logger}
}

func (mw loggingMiddleware) Shorten(s shortUrl) (output shortUrl, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "shorten",
			"full", s.FullUrl,
			"short", s.ShortUrl,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Shorten(s)
	return
}

func (mw loggingMiddleware) Fetch(id string) (s shortUrl, err error) {
	response, err := mw.next.Fetch(id)

	defer func(begin time.Time, response shortUrl) {
		mw.logger.Log(
			"method", "fetch",
			"id", id,
			"took", time.Since(begin),
		)
	}(time.Now(), response)

	return response, err
}
