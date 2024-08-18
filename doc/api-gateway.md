You're right to ask! You don’t need to use Docker to achieve the same result. We can still use Traefik and Consul without Docker by running everything natively on your host machine.

Here’s how you can set up a system where Traefik dynamically routes traffic to two different Go services (live and sandbox) based on the `Authorization` header, all without Docker.

### Setup Overview

1. **Consul**: Used as the service registry.
2. **Traefik**: Acts as the reverse proxy and load balancer, routing traffic based on the `Authorization` header.
3. **Two Go Services**: `go-api-live` and `go-api-sandbox`, each with different logic and configurations.

### Step-by-Step Setup

#### 1. **Install Traefik**

First, download and install Traefik:

```bash
curl -LO https://github.com/traefik/traefik/releases/download/v2.10.4/traefik_v2.10.4_linux_amd64.tar.gz
tar -xvf traefik_v2.10.4_linux_amd64.tar.gz
sudo mv traefik /usr/local/bin/
```

#### 2. **Start Consul**

Make sure Consul is running locally:

```bash
consul agent -dev -client=0.0.0.0
```

You can access the Consul UI at [http://localhost:8500](http://localhost:8500).

#### 3. **Configure Traefik**

Create a Traefik configuration file (`traefik.yml`) to integrate with Consul and handle routing based on the `Authorization` header.

```yaml
entryPoints:
  web:
    address: ":80"

providers:
  consulCatalog:
    endpoint:
      address: "localhost:8500"
    exposedByDefault: false

api:
  dashboard: true
  insecure: true # Enables the Traefik dashboard at http://localhost:8080
```

This configuration tells Traefik to connect to Consul and look for services registered with it. It also enables the dashboard at `http://localhost:8080`.

#### 4. **Implement the Go Services**

We’ll create two Go applications: one for live payments and one for sandbox payments.

##### `go-api-live/main.go`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hashicorp/consul/api"
)

func registerServiceWithConsul() {
	// Create a new Consul client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "localhost:8500"
	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Service registration details
	registration := &api.AgentServiceRegistration{
		ID:      "go-api-live",
		Name:    "go-api",
		Port:    8081,
		Address: "localhost",
		Tags:    []string{"traefik.http.routers.go-api.rule=HeadersRegexp(`Authorization`, `^sk_live_|pk_live_`)"},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://localhost:8081/health",
			Interval: "10s",
		},
	}

	// Register the service with Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	fmt.Println("Live service registered with Consul")
}

func main() {
	// Register the service with Consul
	registerServiceWithConsul()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Live instance is healthy"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handling live payments"))
	})

	log.Println("Starting live server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
```

##### `go-api-sandbox/main.go`

The sandbox version has similar code but uses a different routing rule:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hashicorp/consul/api"
)

func registerServiceWithConsul() {
	// Create a new Consul client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "localhost:8500"
	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Service registration details
	registration := &api.AgentServiceRegistration{
		ID:      "go-api-sandbox",
		Name:    "go-api",
		Port:    8082,
		Address: "localhost",
		Tags:    []string{"traefik.http.routers.go-api.rule=HeadersRegexp(`Authorization`, `^sk_sandbox_|pk_sandbox_`)"},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://localhost:8082/health",
			Interval: "10s",
		},
	}

	// Register the service with Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	fmt.Println("Sandbox service registered with Consul")
}

func main() {
	// Register the service with Consul
	registerServiceWithConsul()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sandbox instance is healthy"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handling sandbox payments"))
	})

	log.Println("Starting sandbox server on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
```

#### 5. **Running the Services**

1. Start Consul:

```bash
consul agent -dev -client=0.0.0.0
```

2. Run the live service:

```bash
go run go-api-live/main.go
```

3. Run the sandbox service:

```bash
go run go-api-sandbox/main.go
```

#### 6. **Configure Hostnames in `/etc/hosts`**

To route requests using custom hostnames like `api.host.local`, update your `/etc/hosts` file:

```bash
sudo nano /etc/hosts
```

Add:

```
127.0.0.1 api.host.local
```

#### 7. **Start Traefik**

Run Traefik with your configuration file:

```bash
traefik --configFile=traefik.yml
```

#### 8. **Testing with Postman**

Now you can test your setup with Postman by setting the `Authorization` header:

- **For live payments**:
  - `Authorization: sk_live_xxxxx` or `Authorization: pk_live_xxxxx`
  - The request is routed to the live service (`go-api-live`).

- **For sandbox payments**:
  - `Authorization: sk_sandbox_xxxxx` or `Authorization: pk_sandbox_xxxxx`
  - The request is routed to the sandbox service (`go-api-sandbox`).

Both services are accessible via the same base URL: `http://api.host.local`, but routing is determined by the `Authorization` header.

### Conclusion

This setup demonstrates how you can use Traefik, Consul, and Go applications without Docker to dynamically route traffic based on the `Authorization` header. This is a powerful pattern for managing multiple service environments (e.g., live and sandbox) while providing a unified API interface to your clients.
