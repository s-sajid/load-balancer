package main

import "net/http/httputil"

type server struct {
	addr  string
	proxy *httputil.ReverseProxy
}
