package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func (dc *DoitClient) createAPIURL(kind string) string {
	fURL := strings.Join([]string{kind}, "/")
	return fURL
}

func (dc *DoitClient) GetHost(*Host, error) {
	res, err := http.Get(dc.createAPIURL("host"))
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
