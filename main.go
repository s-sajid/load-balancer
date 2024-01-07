package main

import (
	"fmt"
	"sync"

	"github.com/s-sajid/loadbalancer/loadbalancer"
	"github.com/s-sajid/loadbalancer/servers"
)

func main() {
	var wg sync.WaitGroup

	numServers := 5
	startPort := 8080

	// Start servers
	wg.Add(1)
	go func() {
		defer wg.Done()
		servers.RunServers(numServers)
	}()

	fmt.Printf("There are %d servers running live ports %d to %d\n", numServers, startPort, startPort+numServers-1)

	// Start load balancer
	wg.Add(1)
	go func() {
		defer wg.Done()
		loadbalancer.CreateLoadBalancer(numServers)
	}()

	wg.Wait()
}
