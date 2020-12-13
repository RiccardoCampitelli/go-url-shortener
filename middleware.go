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

func (mw loggingMiddleware) Shorten(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "shorten",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Shorten(s)
	return
}

func (mw loggingMiddleware) Fetch(s string) (n string) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "fetch",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Fetch(s)
	return
}
