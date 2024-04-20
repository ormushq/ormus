# Docker compose

This provide local development environment.

## Requirement

- docker
- docker-compose

## Usage

for use first copy ``serve/docker-compose/.env.compose.example`` to ``serve/docker-compose/.env`` than
run bellow command 

```shell
 .\serve\docker-compose.bat up
```

in linux os use this command

```shell
 ./serve/docker-compose.bash up
```

## Specification

All containers list:

- destination (use configurable base image from env file)
- redis (use redis:alpine image)
- rabbitmq (use rabbitmq:3-management-alpine image)
- traefik (use traefik:latest image)

## Rabbitmq

You can access to rabbitmq management with the address bellow:

``
rabbitmq.${COMPOSE_PROJECT_NAME}.local
``

note:

- The COMPOSE_PROJECT_NAME parameter is configured in .env file
- You must define host first in your host file with IP 127.0.0.1

