package vex_go

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	apiToken string
	host     string
	scheme   string
	client   *http.Client
	debug    bool
}

type Response struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func New(apiToken string, options ...ClientOption) (*Client, error) {
	client := &Client{
		apiToken: apiToken,
		host:     "vex.broswen.com",
		scheme:   "https://",
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
	for _, option := range options {
		option(client)
	}
	return client, nil
}

type ClientOption func(client *Client)

func WithHost(host string) ClientOption {
	return func(client *Client) {
		client.host = host
	}
}

func WithScheme(scheme string) ClientOption {
	return func(client *Client) {
		client.scheme = scheme
	}
}

func WithDebug(debug bool) ClientOption {
	return func(client *Client) {
		client.debug = debug
	}
}

func (c *Client) doRequestContext(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.scheme+c.host+url, body)
	if err != nil {
		return nil, err
	}
	if c.debug {
		log.Printf("%v %v", method, url)
		log.Printf("%#v", req)
	}
	req.Header.Add("Authorization", "Basic "+c.apiToken)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("error: %v", resp.Status)
	}
	if c.debug {
		log.Printf("%#v", resp)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if c.debug {
		log.Printf("%#v", string(respBody))
	}
	return respBody, err
}
