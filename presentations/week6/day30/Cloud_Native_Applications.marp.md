---
marp: true
theme: default
paginate: true
---

# Cloud-Native Go Applications
## Day 30: Building for the Cloud

---

# Overview

1. Cloud-Native Principles
2. Containerization
3. Kubernetes Integration
4. Service Discovery
5. Configuration Management
6. Observability

---

# Cloud-Native Principles

1. Microservices
2. Containerization
3. Dynamic Management
4. Automation
5. Scalability

---

# Containerization with Go

1. Multi-stage Builds
```dockerfile
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest
COPY --from=builder /app/app .
CMD ["./app"]
```

2. Best Practices
- Small base images
- Layer optimization
- Security considerations

---

# Kubernetes Integration

1. Client-go Usage
```go
config, err := rest.InClusterConfig()
clientset, err := kubernetes.NewForConfig(config)
```

2. Custom Controllers
```go
type Controller struct {
    informerFactory informers.SharedInformerFactory
    serviceLister   listers.ServiceLister
    serviceSynced   cache.InformerSynced
}
```

---

# Service Discovery

1. DNS-based
```go
type ServiceDiscovery struct {
    Namespace string
    Service   string
    Port      int
}
```

2. Kubernetes API
```go
func (sd *ServiceDiscovery) GetEndpoints() ([]string, error)
```

---

# Configuration Management

1. Environment Variables
```go
type Config struct {
    DatabaseURL string `env:"DB_URL,required"`
    Port        int    `env:"PORT" envDefault:"8080"`
}
```

2. ConfigMaps
```go
func loadConfigMap(name string) (map[string]string, error)
```

---

# Health Checks

1. Liveness Probe
```go
func livenessHandler(w http.ResponseWriter, r *http.Request) {
    // Check core functionality
}
```

2. Readiness Probe
```go
func readinessHandler(w http.ResponseWriter, r *http.Request) {
    // Check external dependencies
}
```

---

# Observability

1. Metrics
```go
func initPrometheus() {
    // Register custom metrics
}
```

2. Distributed Tracing
```go
func initTracing() {
    // Set up OpenTelemetry
}
```

---

# Scaling Strategies

1. Horizontal Pod Autoscaling
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: myapp
spec:
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50
```

2. Custom Metrics
```go
type CustomMetrics struct {
    QueueLength   int
    ResponseTime  time.Duration
}
```

---

# Resource Management

1. Resource Requests
```yaml
resources:
  requests:
    memory: "64Mi"
    cpu: "250m"
  limits:
    memory: "128Mi"
    cpu: "500m"
```

2. Go Runtime Configuration
```go
func tuneRuntime() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}
```

---

# Resilience Patterns

1. Circuit Breaker
2. Retries with Backoff
3. Rate Limiting
4. Fallback Strategies
5. Bulkhead Pattern

---

# Deployment Strategies

1. Rolling Updates
2. Blue-Green Deployment
3. Canary Releases
4. A/B Testing
5. Feature Flags

---

# Security Best Practices

1. Image Security
2. Network Policies
3. RBAC Configuration
4. Secret Management
5. Security Context

---

# Hands-on Exercise

Build a cloud-native application with:
1. Kubernetes deployment
2. Health checks
3. Metrics collection
4. Auto-scaling
5. Proper configuration

---

# Key Takeaways

1. Design for failure
2. Implement proper monitoring
3. Use container best practices
4. Manage configuration properly
5. Consider scalability

---

# Next Steps

1. Advanced Kubernetes Patterns
2. Custom Controllers
3. Service Mesh Integration
4. GitOps Workflows
5. Advanced Monitoring
