package vex_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Account struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	ModifiedOn  time.Time `json:"modified_on"`
}

type AccountResponse struct {
	Data Account `json:"data"`
	Response
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	url := fmt.Sprintf("/accounts/%s", id)
	resp, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	a := &AccountResponse{}
	err = json.Unmarshal(resp, a)
	if err != nil {
		return nil, err
	}
	return &a.Data, nil
}

func (c *Client) UpdateAccount(ctx context.Context, a *Account) error {
	url := fmt.Sprintf("/accounts/%s", a.ID)
	j, err := json.Marshal(a)
	if err != nil {
		return err
	}
	resp, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	ar := &AccountResponse{}
	err = json.Unmarshal(resp, a)
	if err != nil {
		return err
	}
	*a = ar.Data
	return nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string) error {
	url := fmt.Sprintf("/accounts/%s", id)
	_, err := c.doRequestContext(ctx, http.MethodDelete, url, nil)
	return err
}
