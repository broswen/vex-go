package vex_go

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var project = &Project{
	ID:          "abc123",
	AccountID:   "abc123",
	Name:        "test project",
	Description: "test project",
	CreatedOn:   someTime,
	ModifiedOn:  someTime,
}

var projects = []*Project{
	{
		ID:          "abc123",
		AccountID:   "abc123",
		Name:        "test project",
		Description: "test project",
		CreatedOn:   someTime,
		ModifiedOn:  someTime,
	},
}
var projectJSON = fmt.Sprintf(`
{
	"data": {
			"id": "abc123",
			"account_id": "abc123",
			"name": "test project",
			"description": "test project",
			"created_on": "%s",
			"modified_on": "%s"
	},
	"success": true,
	"errors": []
}
`, someTime.Format(time.RFC3339), someTime.Format(time.RFC3339))

var projectsJSON = fmt.Sprintf(`
{
	"data": [{
			"id": "abc123",
			"account_id": "abc123",
			"name": "test project",
			"description": "test project",
			"created_on": "%s",
			"modified_on": "%s"
	}],
	"success": true,
	"errors": []
}
`, someTime.Format(time.RFC3339), someTime.Format(time.RFC3339))

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
		fmt.Fprintf(w, projectsJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects", handler)
	want := projects
	actual, err := client.GetProjects(context.Background(), "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)

	}
}
