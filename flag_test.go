package vex_go

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var flag = &Flag{
	ID:        "abc123",
	ProjectID: "abc123",
	AccountID: "abc123",
	Key:       "feature1",
	Type:      "BOOLEAN",
	Value:     "false",
}
var flagJSON = `
{
	"id": "abc123",
	"account_id": "abc123",
	"project_id": "abc123",
	"key": "feature1",
	"type": "BOOLEAN",
	"value": "false"
}
`

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
		fmt.Fprintf(w, fmt.Sprintf("[%s]", flagJSON))
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123/projects/abc123/flags", handler)
	want := []*Flag{flag}
	actual, err := client.GetFlags(context.Background(), "abc123", "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
