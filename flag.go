package vex_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FlagType string

const (
	STRING  FlagType = "STRING"
	BOOLEAN FlagType = "BOOLEAN"
	NUMBER  FlagType = "NUMBER"
)

type Flag struct {
	ID         string    `json:"id"`
	ProjectID  string    `json:"project_id"`
	AccountID  string    `json:"account_id"`
	Key        string    `json:"key"`
	Type       FlagType  `json:"type"`
	Value      string    `json:"value"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}

func (f Flag) ToJSON() ([]byte, error) {
	return json.Marshal(&f)
}

type FlagResponse struct {
	Data Flag `json:"data"`
	Response
}

type ListFlagResponse struct {
	Data []*Flag `json:"data"`
	Response
}

func (c *Client) CreateFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags", f.AccountID, f.ProjectID)
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	resp, err := c.doRequestContext(ctx, http.MethodPost, url, bytes.NewReader(j))
	if err != nil {
		return err
	}

	fr := &FlagResponse{}
	err = json.Unmarshal(resp, fr)
	if err != nil {
		return err
	}
	*f = fr.Data
	return nil
}

func (c *Client) GetFlags(ctx context.Context, accountId, projectId string) ([]*Flag, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags", accountId, projectId)
	resp, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	lf := &ListFlagResponse{}
	err = json.Unmarshal(resp, lf)
	if err != nil {
		return nil, err
	}
	return lf.Data, nil
}

func (c *Client) GetFlag(ctx context.Context, accountId, projectId, flagId string) (*Flag, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", accountId, projectId, flagId)
	resp, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	fr := &FlagResponse{}
	err = json.Unmarshal(resp, fr)
	if err != nil {
		return nil, err
	}
	return &fr.Data, nil
}

func (c *Client) UpdateFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", f.AccountID, f.ProjectID, f.ID)
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	resp, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	fr := &FlagResponse{}
	err = json.Unmarshal(resp, fr)
	if err != nil {
		return err
	}
	*f = fr.Data
	return nil
}

func (c *Client) DeleteFlag(ctx context.Context, f *Flag) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s/flags/%s", f.AccountID, f.ProjectID, f.ID)
	_, err := c.doRequestContext(ctx, http.MethodDelete, url, nil)
	return err
}
