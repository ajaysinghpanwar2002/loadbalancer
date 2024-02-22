package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(name, urlStr string) *server {
	urlParsed, err := url.Parse(urlStr)
	if err != nil {
		panic(fmt.Sprintf("failed to parse server URL '%s': %v", urlStr, err))
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(urlParsed)
	return &server{
		Name:         name,
		URL:          urlStr,
		ReverseProxy: reverseProxy,
		Health:       true, // Assume server is healthy initially
	}
}

func (s *server) checkHealth(timeout time.Duration) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Head(s.URL) // Simple HEAD request to check if the server is up
	mutex.Lock()                    // Ensure thread-safe access to the Health attribute
	defer mutex.Unlock()
	if err != nil || resp.StatusCode != http.StatusOK {
		s.Health = false
	} else {
		s.Health = true
	}
}
