package vex_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Project struct {
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Client) CreateProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}
	data, err := c.doRequestContext(ctx, http.MethodPost, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(data)).Decode(p)
}

func (c *Client) GetProjects(ctx context.Context, accountId string) ([]*Project, error) {
	url := fmt.Sprintf("/accounts/%s/projects", accountId)
	data, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var projects []*Project
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetProject(ctx context.Context, accountId, projectId string) (*Project, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s", accountId, projectId)
	data, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var p Project
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *Client) UpdateProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}
	data, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(data)).Decode(p)
}

func (c *Client) DeleteProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	_, err := c.doRequestContext(ctx, http.MethodDelete, url, nil)
	return err
}
