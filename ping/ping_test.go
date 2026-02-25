package main

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// func TestPing(t *testing.T) {
// 	client := &http.Client{}
// 	pinger := Pinger{client}
// 	got := pinger.Ping("https://example.com")

// 	if !got {
// 		t.Errorf("Expected example.com to be available")
// 	}

// 	got = pinger.Ping("https://example.com/404")

// 	if got {
// 		t.Errorf("Expected example.com/404 to be unavailable")
// 	}
// }

type MockClient struct{}

func (c *MockClient) Head(url string) (resp *http.Response, err error) {
	parts := strings.Split(url, "/")
	last := parts[len(parts)-1]
	statusCode, err := strconv.Atoi(last)

	if err != nil {
		return nil, err
	}

	resp = &http.Response{StatusCode: statusCode}
	return resp, nil
}

func TestPing(t *testing.T) {
	client := &MockClient{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com/200")

	if !got {
		t.Errorf("Expected example.com/200 to be available")
	}

	got = pinger.Ping("https://example.com/404")

	if got {
		t.Errorf("Expected example.com/404 to be unavailable")
	}
}
