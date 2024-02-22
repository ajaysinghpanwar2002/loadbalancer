package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LoadBalancer interface {
	SelectServer() (*server, error)
}

type RoundRobinLoadBalancer struct{}
type RandomLoadBalancer struct{}
type LeastConnectionsLoadBalancer struct{} // Placeholder

func (r *RoundRobinLoadBalancer) SelectServer() (*server, error) {
	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < len(serverList); i++ {
		nextIndex := (lastServedIndex + 1) % len(serverList)
		server := serverList[nextIndex]
		lastServedIndex = nextIndex
		if server.Health {
			return server, nil
		}
	}
	return nil, fmt.Errorf("no healthy server found in round robin")
}

func (r *RandomLoadBalancer) SelectServer() (*server, error) {
	rand.Seed(time.Now().UnixNano())
	mutex.Lock()
	defer mutex.Unlock()
	healthyServers := []*server{}
	for _, server := range serverList {
		if server.Health {
			healthyServers = append(healthyServers, server)
		}
	}
	if len(healthyServers) == 0 {
		return nil, fmt.Errorf("no healthy server found in random selection")
	}
	return healthyServers[rand.Intn(len(healthyServers))], nil
}

func (l *LeastConnectionsLoadBalancer) SelectServer() (*server, error) {
	// Placeholder for real implementation
	return nil, fmt.Errorf("least connections strategy not implemented")
}
