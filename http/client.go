// Copyright (C) 2021 David Bundgaard
package http

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Region string

const (
	Europe = "eu"
	USA    = "us"
)

type Client struct {
	httpClient *http.Client

	credentials credentials
	token       string
	apiURL      string
	oauthURL    string
}

type credentials struct {
	ID, Secret string
}

type oauthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c credentials) basicAuth() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ID, c.Secret))))
}

func New(clientId, clientSecret, region string) *Client {
	if clientId == "" || clientSecret == "" {
		log.Fatalf("client ID or client Secret cannot be empty")
	}

	apiURL := fmt.Sprintf("https://%s.api.blizzard.com", region) // api
	oauthURL := fmt.Sprintf("https://%s.battle.net/", region)    // oauth

	return &Client{
		httpClient:  http.DefaultClient,
		credentials: credentials{ID: clientId, Secret: clientSecret},
		apiURL:      apiURL,
		oauthURL:    oauthURL,
	}
}

func (c Client) NewRequest(method string, body io.Reader, format string, v ...interface{}) (*http.Request, error) {
	url := fmt.Sprintf(c.apiURL+format, v...)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func (c Client) TokenRequest(method string, body io.Reader, format string, v ...interface{}) (*http.Response, error) {
	req, err := c.NewRequest(method, body, format, v...)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request %s %d %s", format, resp.StatusCode, resp.Status)
	}
	return resp, nil
}

func (c *Client) Login() error {
	formData := url.Values{}
	formData.Add("grant_type", "client_credentials")
	req, _ := http.NewRequest("POST", c.oauthURL+"/oauth/token", strings.NewReader(formData.Encode()))
	req.Header.Set("Authorization", c.credentials.basicAuth())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	log.Printf("%d %s", resp.StatusCode, resp.Status)
	var t oauthToken
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return err
	}
	c.token = t.AccessToken
	return nil

}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
