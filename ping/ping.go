package main

import (
	"net/http"
)

type Pinger struct {
	client HTTPClient
}

type HTTPClient interface {
	Head(s string) (resp *http.Response, err error)
}

func (p *Pinger) Ping(url string) bool {
	resp, err := p.client.Head(url)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	return true
}
