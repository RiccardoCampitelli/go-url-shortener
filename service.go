package main

import (
	"errors"
)

type UrlShortenerService interface {
	Shorten(string) (string, error)
	Fetch(string) string
}

type urlShortenerService struct {
	repository Repository
}

func (svc urlShortenerService) Shorten(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	svc.repository.FindById("123")

	return "shorter", nil
}

func (urlShortenerService) Fetch(s string) string {
	return "thestring"
}

var ErrEmpty = errors.New("Empty string")
