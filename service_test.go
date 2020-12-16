package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceShorten(t *testing.T) {
	svc := makeMockService()

	testShortUrl := shortUrl{FullUrl: "www.test.com", ShortUrl: "tst"}

	res, err := svc.Shorten(testShortUrl)

	assert.Equal(t, err, nil)
	assert.Equal(t, res, testShortUrl)

}

func TestServiceFetch(t *testing.T) {
	svc := makeMockService()

	id := "123"

	res, err := svc.Fetch(id)

	assert.Equal(t, err, nil)
	assert.Equal(t, res.FullUrl, "www.fromPersistance.com")
	assert.Equal(t, res.ShortUrl, "fps")
}

func makeMockService() UrlShortenerService {

	var svc UrlShortenerService

	mockRepository := mockRepo{}

	svc = urlShortenerService{repository: mockRepository}

	return svc
}

type mockRepo struct{}

func (mockRepo) Init() error {
	return nil
}

func (mockRepo) FindById(id string) (shortUrl, error) {

	su := shortUrl{FullUrl: "www.fromPersistance.com", ShortUrl: "fps"}

	return su, nil

}

func (mockRepo) InsertOne(su shortUrl) error {
	return nil
}
