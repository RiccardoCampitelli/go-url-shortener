package main

type shortenUrlRequest struct {
	S string `json:"s"`
}

type shortenUrlResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type findUrlRequest struct {
	S string `json:"s"`
}

type findUrlResponse struct {
	V string `json:"v"`
}
