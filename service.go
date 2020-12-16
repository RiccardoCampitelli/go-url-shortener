package main

import (
	"errors"
)

type UrlShortenerService interface {
	Shorten(shortUrl) (shortUrl, error)
	Fetch(string) (shortUrl, error)
}

type urlShortenerService struct {
	repository Repository
}

func (svc urlShortenerService) Shorten(s shortUrl) (shortUrl, error) {

	err := svc.repository.InsertOne(s)

	return s, err
}

func (svc urlShortenerService) Fetch(s string) (shortUrl, error) {

	url, err := svc.repository.FindById(s)

	if err != nil {
		return url, err
	}

	return url, nil
}

var ErrEmpty = errors.New("Empty string")
