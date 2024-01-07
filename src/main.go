package main

import (
	"net/http/httputil"
	"net/url"
)

type server struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newServer(addr string) *server {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &server{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
