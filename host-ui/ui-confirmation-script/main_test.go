package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"testing"
)

func TestHTTPSConnection(t *testing.T) {
    url := "https://www.example.com" // Replace with your domain

    // Skip certificate verification if using a self-signed certificate
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    resp, err := http.Get(url)
    if err != nil {
        t.Fatalf("Failed to make HTTPS request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    expectedContent := `<html>
    <head>
    <title>Hello World</title>
    </head>
    <body>
    <h1>Hello World!</h1>
    </body>
    </html>`

    if string(body) != expectedContent {
        t.Errorf("Content does not match expected content")
    }
}

func TestHTTPRedirect(t *testing.T) {
    url := "http://www.example.com" // Replace with your domain

    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    resp, err := client.Get(url)
    if err != nil {
        t.Fatalf("Failed to make HTTP request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusMovedPermanently && resp.StatusCode != http.StatusFound {
        t.Errorf("Expected redirect status code, got %d", resp.StatusCode)
    }

    location := resp.Header.Get("Location")
    if location != "https://www.example.com/" {
        t.Errorf("Expected redirect to HTTPS, got %s", location)
    }
}