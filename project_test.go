package vex_go

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var project = &Project{
	ID:          "abc123",
	AccountID:   "abc123",
	Name:        "test project",
	Description: "test project",
}
var projectJSON = `
{
	"id": "abc123",
	"account_id": "abc123",
	"name": "test project",
	"description": "test project"
}
`

func TestGetProject(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, projectJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects/abc123", handler)
	want := project
	actual, err := client.GetProject(context.Background(), "abc123", "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetProjects(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf("[%s]", projectJSON))
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects", handler)
	want := []*Project{project}
	actual, err := client.GetProjects(context.Background(), "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)

	}
}
