# Deployment

This guide will walk you through the steps to deploy the Golang microservice application using Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone the repository and navigate to the root directory of the project:

```bash
git clone https://github.com/infranyx/go-grpc-template
cd go-grpc-template
```

2. Create a .env file in the `envs` directory of the project and copy the `local.env` environment variables.

3. Run the following command to build and start the containers:

```bash
docker-compose up -d --build
```

This will build and start the following containers:

- `postgres`: PostgreSQL database
- `kafka`: Kafka message broker
- `zookeeper`: In the context of Kafka, Zookeeper is used to store metadata about the Kafka cluster and its topics. It helps the Kafka brokers maintain their cluster membership and elect leaders, and it also helps clients discover the location of the Kafka brokers.
- `redis`: Redis cache
- `sentry`: Sentry error tracking service
- `app`: Golang microservice application
