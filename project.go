package vex_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Project struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	ModifiedOn  time.Time `json:"modified_on"`
}

type ProjectResponse struct {
	Data Project `json:"data"`
	Response
}

type ListProjectResponse struct {
	Data []*Project `json:"data"`
	Response
}

func (c *Client) CreateProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}
	resp, err := c.doRequestContext(ctx, http.MethodPost, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	pr := &ProjectResponse{}
	err = json.Unmarshal(resp, p)
	if err != nil {
		return err
	}
	*p = pr.Data
	return nil
}

func (c *Client) GetProjects(ctx context.Context, accountId string) ([]*Project, error) {
	url := fmt.Sprintf("/accounts/%s/projects", accountId)
	resp, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	p := &ListProjectResponse{}
	err = json.Unmarshal(resp, p)
	if err != nil {
		return nil, err
	}
	return p.Data, nil
}

func (c *Client) GetProject(ctx context.Context, accountId, projectId string) (*Project, error) {
	url := fmt.Sprintf("/accounts/%s/projects/%s", accountId, projectId)
	resp, err := c.doRequestContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	p := &ProjectResponse{}
	err = json.Unmarshal(resp, p)
	if err != nil {
		return nil, err
	}
	return &p.Data, nil
}

func (c *Client) UpdateProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}
	resp, err := c.doRequestContext(ctx, http.MethodPut, url, bytes.NewReader(j))
	if err != nil {
		return err
	}
	pr := &ProjectResponse{}
	err = json.Unmarshal(resp, p)
	if err != nil {
		return err
	}
	*p = pr.Data
	return nil
}

func (c *Client) DeleteProject(ctx context.Context, p *Project) error {
	url := fmt.Sprintf("/accounts/%s/projects/%s", p.AccountID, p.ID)
	_, err := c.doRequestContext(ctx, http.MethodDelete, url, nil)
	return err
}
