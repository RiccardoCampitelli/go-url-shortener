package main

import (
	"errors"
	"fmt"
)

type UrlShortenerService interface {
	Shorten(shortUrl) (string, error)
	Fetch(string) (shortUrl, error)
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

func (svc urlShortenerService) Fetch(s string) (shortUrl, error) {

	fmt.Println(s)

	url, err := svc.repository.FindById(s)

	if err != nil {
		return url, err
	}

	return url, nil
}

var ErrEmpty = errors.New("Empty string")
