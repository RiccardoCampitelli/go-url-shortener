package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	var svc UrlShortenerService

	repo := &repository{}

	err := repo.Init()

	if err != nil {
		fmt.Println(err)
	}

	svc = urlShortenerService{repository: repo}

	router := mux.NewRouter()

	middleware := loggingMiddleware{}

	svc = middleware.Wrap(svc)

	shortenUrlHandler := httptransport.NewServer(
		makeShortenUrlEndpoint(svc),
		decodeShortenUrlRequest,
		encodeResponse,
	)

	findUrlHandler := httptransport.NewServer(
		makeFindUrlEndpoint(svc),
		decodeFindUrlRequest,
		encodeResponse,
	)

	router.Methods("GET").Path("/").Handler(findUrlHandler)

	router.Methods("GET").Path("/shorten").Handler(shortenUrlHandler)

	http.ListenAndServe(":8080", router)
}

func decodeShortenUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request shortenUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeFindUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request findUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
