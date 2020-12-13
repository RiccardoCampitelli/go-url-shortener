package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeShortenUrlEndpoint(svc UrlShortenerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(shortenUrlRequest)
		v, err := svc.Shorten(req.S)
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
