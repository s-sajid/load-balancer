package main

import (
	"sync"

	"github.com/s-sajid/loadbalancer/loadbalancer"
	"github.com/s-sajid/loadbalancer/servers"
)

func main() {
	var wg sync.WaitGroup

	numServers := 5

	// Start servers
	wg.Add(1)
	go func() {
		defer wg.Done()
		servers.RunServers(numServers)
	}()

	// Start load balancer
	wg.Add(1)
	go func() {
		defer wg.Done()
		loadbalancer.CreateLoadBalancer(numServers)
	}()

	wg.Wait()
}
