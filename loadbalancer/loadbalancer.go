package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

var (
	baseURL = "http://localhost:808"
)

type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

type Endpoints struct {
	List []*url.URL
}

func (e *Endpoints) Cycle() {
	temp := e.List[0]
	e.List = e.List[1:]
	e.List = append(e.List, temp)
}

func CreateLoadBalancer(num int) {
	var lb LoadBalancer
	var ep Endpoints

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	for i := 0; i < num; i++ {
		ep.List = append(ep.List, createEnpoint(baseURL, i))
	}

	router.HandleFunc("/loadbalancer", makeRequest(&lb, &ep))

	fmt.Println("Load Balancer is up and running on port:", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func makeRequest(lb *LoadBalancer, ep *Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for !testServer(ep.List[0].String()) {
			ep.Cycle()
		}

		lb.RevProxy = *httputil.NewSingleHostReverseProxy(ep.List[0])
		ep.Cycle()
		lb.RevProxy.ServeHTTP(w, r)
	}
}

func createEnpoint(endpoint string, index int) *url.URL {
	link := endpoint + strconv.Itoa(index)
	url, err := url.Parse(link)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return url
}

func testServer(endpoint string) bool {
	resp, err := http.Get(endpoint)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
