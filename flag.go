package vex_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlagType string

const (
	STRING  FlagType = "STRING"
	BOOLEAN FlagType = "BOOLEAN"
	NUMBER  FlagType = "NUMBER"
)

type Flag struct {
	ID        string   `json:"id"`
	ProjectID string   `json:"project_id"`
	AccountID string   `json:"account_id"`
	Key       string   `json:"key"`
	Type      FlagType `json:"type"`
	Value     string   `json:"value"`
}

func (f Flag) ToJSON() ([]byte, error) {
	return json.Marshal(&f)
}

func (c *Client) CreateFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags", f.AccountID, f.ProjectID)
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	data, err := c.doRequestContext(ctx, http.MethodPost, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(data)).Decode(f)
}

func (c *Client) GetFlags(ctx context.Context, accountId, projectId string) ([]*Flag, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags", accountId, projectId)
	data, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var flags []*Flag
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&flags)
	if err != nil {
		return nil, err
	}
	return flags, nil
}

func (c *Client) GetFlag(ctx context.Context, accountId, projectId, flagId string) (*Flag, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", accountId, projectId, flagId)
	data, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var f Flag
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (c *Client) UpdateFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", f.AccountID, f.ProjectID, f.ID)
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	data, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(data)).Decode(f)
}

func (c *Client) DeleteFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", f.AccountID, f.ProjectID, f.ID)
	_, err := c.doRequestContext(ctx, http.MethodDelete, url, nil)
	return err
}
