# Cloud-Native Go Application

This demo implements a cloud-native Go application showcasing various patterns and best practices.

## Features

1. Cloud-Native Characteristics
   - Health checks
   - Metrics (Prometheus)
   - Tracing (OpenTelemetry)
   - Graceful shutdown
   - Configuration management

2. Monitoring & Observability
   - Request metrics
   - Latency tracking
   - Health status
   - Readiness probes

3. Production-Ready Features
   - Middleware support
   - Error handling
   - Graceful shutdown
   - Connection management

## Configuration

The application supports configuration through environment variables:

```bash
PORT=8080
LOG_LEVEL=info
```

## Metrics

Prometheus metrics available at `/metrics`:
- http_requests_total
- http_request_duration_seconds

## Health Checks

1. Liveness Probe: `/health`
2. Readiness Probe: `/ready`

## API Endpoints

- GET /api/data - Sample data endpoint
- GET /metrics - Prometheus metrics
- GET /health - Health check
- GET /ready - Readiness check

## Running in Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cloud-native-app
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: app
        image: cloud-native-app:latest
        ports:
        - containerPort: 8080
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
```

## Building

```bash
docker build -t cloud-native-app .
```

## Running Locally

```bash
go run main.go
```

## Best Practices Demonstrated

1. Proper Health Checks
   - Liveness probe
   - Readiness probe
   - Detailed status

2. Metrics & Monitoring
   - Prometheus integration
   - Request tracking
   - Latency monitoring

3. Production Readiness
   - Graceful shutdown
   - Error handling
   - Connection management

4. Cloud-Native Design
   - Stateless design
   - Configuration via environment
   - Container-ready

## Requirements

- Go 1.21+
- Prometheus client
- OpenTelemetry
