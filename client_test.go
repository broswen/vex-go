package vex_go

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	var err error
	client, err = New("", WithBaseUrl(server.URL), WithScheme(""))
	if err != nil {
		panic(err)
	}
}

func teardown() {
	server.Close()
}
