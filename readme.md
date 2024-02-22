# Go Load Balancer

This project is a custom load balancer written in Go, designed to distribute incoming HTTP requests across a pool of backend servers. It supports multiple load balancing strategies, including Round Robin, Random, and a placeholder for Least Connections. The load balancer improves application scalability, fault tolerance, and resource utilization by dynamically routing traffic to the healthiest servers available.

## Features

- **Multiple Load Balancing Strategies**: Includes Round Robin and Random strategies out of the box, with support for easily adding new strategies.
- **Dynamic Server Management**: Servers can be added or removed from the pool dynamically, allowing for flexible scalability.
- **Health Checks**: Periodic health checks ensure that traffic is only routed to healthy servers, improving reliability.
- **Concurrency Safe**: Utilizes Go's concurrency features to safely manage server lists and handle incoming requests across multiple goroutines.

## Getting Started

### Prerequisites

- Go 1.15 or later.

### Installation

1. Clone the repository to your local machine:
    ```sh
    git clone https://github.com/ajaysinghpanwar2002/loadbalancer.git
    cd loadbalancer
    ```

2. Initialize the Go module (if not already done):
    ```sh
    go mod init loadbalancer
    ```

### Running the Load Balancer

From the project's root directory, run:

```sh
go run loadbalancer.go server.go strategy.go
