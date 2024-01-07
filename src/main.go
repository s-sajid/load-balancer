package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type ServerHandler interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

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

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []ServerHandler
}

func NewLoadBalancer(port string, servers []ServerHandler) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	servers := []ServerHandler{
		newServer("https://www.google.com/"),
		newServer("https://www.bing.com/"),
		newServer("https://www.yahoo.com/"),
	}

	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Server live on port:", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
