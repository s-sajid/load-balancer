package main

import (
	"sync"

	"github.com/s-sajid/loadbalancer/loadbalancer"
	"github.com/s-sajid/loadbalancer/servers"
)

func main() {
	var wg sync.WaitGroup

	// Start servers
	wg.Add(1)
	go func() {
		defer wg.Done()
		servers.RunServers(5)
	}()

	// Start load balancer
	wg.Add(1)
	go func() {
		defer wg.Done()
		loadbalancer.CreateLoadBalancer(5)
	}()

	wg.Wait()
}
