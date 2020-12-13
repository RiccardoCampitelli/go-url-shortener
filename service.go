package main

import (
	"errors"
	"fmt"
)

type UrlShortenerService interface {
	Shorten(shortUrl) (string, error)
	Fetch(string) string
}

type urlShortenerService struct {
	repository Repository
}

func (svc urlShortenerService) Shorten(s shortUrl) (string, error) {
	// if s == "" {
	// 	return "", ErrEmpty
	// }
	// TODO: Validation

	// res, err := svc.repository.FindById(s.fullUrl)
	fmt.Println(s)

	err := svc.repository.InsertOne(s)

	if err != nil {
		return "", err
	}

	return "thing", err
}

func (urlShortenerService) Fetch(s string) string {
	return "thestring"
}

var ErrEmpty = errors.New("Empty string")
