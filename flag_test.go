package vex_go

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var flag = &Flag{
	ID:         "abc123",
	ProjectID:  "abc123",
	AccountID:  "abc123",
	Key:        "feature1",
	Type:       "BOOLEAN",
	Value:      "false",
	CreatedOn:  someTime,
	ModifiedOn: someTime,
}

var flags = []*Flag{
	{
		ID:         "abc123",
		ProjectID:  "abc123",
		AccountID:  "abc123",
		Key:        "feature1",
		Type:       "BOOLEAN",
		Value:      "false",
		CreatedOn:  someTime,
		ModifiedOn: someTime,
	},
}
var flagJSON = fmt.Sprintf(`
{
	"data": {
			"id": "abc123",
			"account_id": "abc123",
			"project_id": "abc123",
			"key": "feature1",
			"type": "BOOLEAN",
			"value": "false",
			"created_on": "%s",
			"modified_on": "%s"
	},
	"success": true,
	"errors": []
}
`, someTime.Format(time.RFC3339), someTime.Format(time.RFC3339))

var flagsJSON = fmt.Sprintf(`
{
	"data": [{
			"id": "abc123",
			"account_id": "abc123",
			"project_id": "abc123",
			"key": "feature1",
			"type": "BOOLEAN",
			"value": "false",
			"created_on": "%s",
			"modified_on": "%s"
	}],
	"success": true,
	"errors": []
}
`, someTime.Format(time.RFC3339), someTime.Format(time.RFC3339))

func TestGetFlag(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, flagJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects/abc123/flags/abc123", handler)
	want := flag
	actual, err := client.GetFlag(context.Background(), "abc123", "abc123", "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetFlags(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, flagsJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects/abc123/flags", handler)
	want := flags
	actual, err := client.GetFlags(context.Background(), "abc123", "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
