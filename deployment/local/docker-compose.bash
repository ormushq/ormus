#! /bin/bash

docker compose \
--env-file ./deployment/local/.env \
-f ./deployment/local/services/etcd.yml \
-f ./deployment/local/services/prometheus.yml \
-f ./deployment/local/services/jaeger.yml \
-f ./deployment/local/services/destination_publisher.yml \
-f ./deployment/local/services/destination_webhook_worker.yml \
-f ./deployment/local/services/otel_collector.yml \
-f ./deployment/local/services/grafana.yml \
-f ./deployment/local/services/redis.yml \
-f ./deployment/local/services/scylladb.yml \
-f ./deployment/local/services/rabbitmq.yml \
-f ./deployment/local/services/traefik.yml \
 "$@"