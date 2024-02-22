package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	serverList      []*server
	lastServedIndex int
	mutex           sync.Mutex
	strategy        LoadBalancer
)

func main() {
	initializeServers()                     // Initialize serverList with predefined servers
	strategy = selectLoadBalancerStrategy() // Select load balancing strategy based on configuration

	http.HandleFunc("/", forwardRequest)
	go startHealthCheck(30*time.Second, 5*time.Second) // Start health check with 30s interval and 5s timeout
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server, err := strategy.SelectServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func initializeServers() {
	// Initialize your server list here
	serverList = []*server{
		newServer("server-1", "http://127.0.0.1:5001"),
		newServer("server-2", "http://127.0.0.1:5002"),
		// Add more servers as needed
	}
}

func selectLoadBalancerStrategy() LoadBalancer {
	// This could be loaded from a configuration file or environment variable
	var strategyName = "round_robin" // Example, change as needed

	switch strategyName {
	case "random":
		return &RandomLoadBalancer{}
	case "least_connections":
		return &LeastConnectionsLoadBalancer{}
	default:
		return &RoundRobinLoadBalancer{}
	}
}

// AddServer adds a new server to the server list
func AddServer(name, url string) {
	mutex.Lock()
	defer mutex.Unlock()
	server := newServer(name, url)
	serverList = append(serverList, server)
}

// RemoveServer removes a server from the server list by name
func RemoveServer(name string) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, server := range serverList {
		if server.Name == name {
			serverList = append(serverList[:i], serverList[i+1:]...)
			break
		}
	}
}

func startHealthCheck(interval, timeout time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for _, srv := range serverList {
			go srv.checkHealth(timeout) // Asynchronously check the health of each server
		}
	}
}
