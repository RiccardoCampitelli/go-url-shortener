package main

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func makeShortenUrlEndpoint(svc UrlShortenerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(shortenUrlRequest)

		su := shortUrl{FullUrl: req.FullUrl, ShortUrl: req.ShortUrl}

		fmt.Println(req, su)

		v, err := svc.Shorten(su)
		if err != nil {
			return shortenUrlResponse{v, err.Error()}, nil
		}
		return shortenUrlResponse{v, ""}, nil
	}
}

func makeFindUrlEndpoint(svc UrlShortenerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findUrlRequest)
		v := svc.Fetch(req.S)
		return findUrlResponse{v}, nil
	}
}
