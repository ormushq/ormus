# Docker Compose Environment

This repository contains the Docker Compose files and configuration for setting up a development environment using the following services:

- Traefik (Reverse Proxy)
- RabbitMQ
- Redis
- etcd
- ScyllaDB
- Prometheus
- Grafana
- Jaeger
- Otel collector
- Destination-Task-Publisher
- Destination-Webhook-Worker


## Getting Started

1. Install Docker and Docker Compose on your development machine.
2. Clone this repository to your local machine.
3. Update the `hosts` file on your development machine to include the following entry:
   127.0.0.1 ormus.local
4. use `go mod vendor` to prevent download packages after every build for golang applications.
5. Run `./deployment/local/docker-compose.bash up -d` in unix and `.\deployment\local\docker-compose.bat up -d` 
in windows. This will start all the services in the background.


## Service Endpoints

- Traefik Dashboard: [http://ormus.local:8080](http://ormus.local:8080)
- RabbitMQ Dashboard: [http://rabbitmq.ormus.local](http://rabbitmq.ormus.local)
- Prometheus: [http://prometheus.ormus.local](http://prometheus.ormus.local)
- Grafana: [http://grafana.ormus.local](http://grafana.ormus.local)
- Jaeger: [http://jaeger.ormuse.local](http://jaeger.ormuse.local)



# generate fake processed event for destination 

```bash
docker exec -it ormus-destination-task-publisher go run cmd/destination/faker/fake_processed_event_producer.go
```
