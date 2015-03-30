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

func (dc *DoitClient) createAPIURL(kind string, key string, v string, domain string) string {
	var fURL string

	if key != "" {
		fURL = strings.Join([]string{dc.URL.String(), kind, key}, "/")
	} else {
		fURL = strings.Join([]string{dc.URL.String(), kind}, "/")
	}

	if kind == "var" && v != "" {
		fURL = strings.Join([]string{fURL, "value", v}, "/")
	}

	if kind == "host" && v != "" {
		fURL = strings.Join([]string{fURL, "host", v}, "/")
	}

	if domain != "" {
		fURL = strings.Join([]string{fURL, "?domain=", domain}, "")
	}

	return fURL
}
