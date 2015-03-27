package main

import (
	"net/url"
	"strings"
)

const (
	JSONMime = "application/json"
)

type DoitClient struct {
	URL        *url.URL
	APIVersion int
}

func (dc *DoitClient) SetURL(apiurl string) error {
	url, err := url.Parse(apiurl)
	if err != nil {
		return err
	}
	dc.URL = url
	return nil
}

func (dc *DoitClient) createAPIURL(kind string, key string, value string, domain string) string {
	fURL := strings.Join([]string{dc.URL.String(), kind, key}, "/")
	if domain != "" {
		fURL = strings.Join([]string{fURL, "?domain=", domain}, "")
	}
	if value != "" && domain != "" {
		fURL = strings.Join([]string{fURL, "&value=", value}, "")
	} else if value != "" && domain == "" {
		fURL = strings.Join([]string{fURL, "?value=", value}, "")
	}
	return fURL
}
