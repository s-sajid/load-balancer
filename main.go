package main

import (
	"github.com/s-sajid/loadbalancer/loadbalancer"
	"github.com/s-sajid/loadbalancer/servers"
)

func main() {
	servers.RunServers(5)
	loadbalancer.CreateLoadBalancer(5)
}
