# Service Mesh Demo with Go and Istio

This demo showcases the implementation of service mesh patterns using Go microservices and Istio.

## Project Structure

```
service-mesh-demo/
├── services/
│   ├── product-service/
│   │   ├── main.go
│   │   └── Dockerfile
│   └── order-service/
│       ├── main.go
│       └── Dockerfile
├── kubernetes/
│   ├── product-service.yaml
│   ├── order-service.yaml
│   └── istio/
│       ├── virtual-service.yaml
│       ├── destination-rules.yaml
│       └── gateway.yaml
└── README.md
```

## Prerequisites

1. Kubernetes cluster
2. Istio installed
3. kubectl configured
4. Docker installed

## Services Overview

1. Product Service:
   - Product catalog API
   - Version-aware responses
   - Health checks

2. Order Service:
   - Order processing
   - Communication with product service
   - Circuit breaker implementation

## Key Features Demonstrated

1. Traffic Management:
   - Canary deployments
   - Circuit breaking
   - Load balancing

2. Security:
   - mTLS communication
   - Service-to-service authentication

3. Observability:
   - Distributed tracing
   - Metrics collection
   - Logging

## Setup Instructions

1. Build the services:
   ```bash
   docker build -t product-service:v1 ./services/product-service
   docker build -t order-service:v1 ./services/order-service
   ```

2. Deploy to Kubernetes:
   ```bash
   kubectl apply -f kubernetes/product-service.yaml
   kubectl apply -f kubernetes/order-service.yaml
   ```

3. Apply Istio configurations:
   ```bash
   kubectl apply -f kubernetes/istio/
   ```

## Testing the Setup

1. Access the services:
   ```bash
   export GATEWAY_URL=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
   curl http://$GATEWAY_URL/api/products
   ```

2. Test circuit breaker:
   ```bash
   hey -z 20s -q 100 http://$GATEWAY_URL/api/products
   ```

## Monitoring

1. Access Grafana:
   ```bash
   istioctl dashboard grafana
   ```

2. Access Jaeger:
   ```bash
   istioctl dashboard jaeger
   ```

## Patterns Demonstrated

1. Circuit Breaker Pattern
2. Retry Pattern
3. Canary Deployment
4. Service Discovery
5. Load Balancing

## Best Practices

1. Always use proper health checks
2. Implement graceful degradation
3. Use structured logging
4. Monitor service metrics
5. Implement proper security measures
