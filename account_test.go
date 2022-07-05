package vex_go

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var account = &Account{
	ID:          "abc123",
	Name:        "test account",
	Description: "test account",
}
var accountJSON = `
{
	"id": "abc123",
	"name": "test account",
	"description": "test account"
}
`

func TestGetAccount(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, accountJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123", handler)
	want := account
	actual, err := client.GetAccount(context.Background(), "abc123")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateAccount(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, accountJSON)
	}
	setup()
	defer teardown()
	mux.HandleFunc("/accounts/abc123", handler)
	want := account
	err := client.UpdateAccount(context.Background(), account)
	if assert.NoError(t, err) {
		assert.Equal(t, want, account)
	}
}
