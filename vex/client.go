package vex

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	apiToken string
	host     string
	client   *http.Client
	debug    bool
}

func New(apiToken string, options ...ClientOption) (*Client, error) {
	client := &Client{
		apiToken: apiToken,
		client:   &http.Client{},
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

func WithDebug(debug bool) ClientOption {
	return func(client *Client) {
		client.debug = debug
	}
}

func (c *Client) doRequestContext(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.host+url, body)
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
