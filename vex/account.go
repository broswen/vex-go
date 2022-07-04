package vex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	url := fmt.Sprintf("/accounts/%s", id)
	data, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	a := &Account{}
	err = json.NewDecoder(bytes.NewReader(data)).Decode(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Client) UpdateAccount(ctx context.Context, a *Account) error {
	url := fmt.Sprintf("/accounts/%s", a.ID)
	j, err := json.Marshal(a)
	if err != nil {
		return err
	}
	data, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(data)).Decode(a)
}
