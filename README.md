# Go Microservice Template

<!-- # Hands-On Microservices with Node.js -->

This is the code repository for [Go Microservice Template](https://www.github.com/infranyx/go-grpc-template), published by InfraNyx.

**Build, test, and deploy robust golang microservices**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Golang Template Project](#golang-template-project)
  - [About the project](#about-the-project)
  - [Prerequisites](#prerequisites)
    - [Design](#design)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Database](#database)
    - [Protocol Buffer](#protocol-buffer)
    - [API docs](#api-docs)
    - [Docker Compose](#docker-compose)
    - [Linting](#linting)
  - [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Golang Template Project

## About the project

The template is used to create golang project. All golang projects must follow the conventions in the
template. Calling for exceptions must be brought up in the engineering team.

## Prerequisites

- Go 1.13+

### Design

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### Status

The template project is in alpha status.

### See also

- [gRPC](https://grpc.io/) for communication between services
- [SQLx](https://github.com/jmoiron/sqlx) for database access and migrations
- [Redis](github.com/go-redis/redis) for caching and message queues
- [Kafka](https://github.com/segmentio/kafka-go) for streaming data processing
- [Echo](https://echo.labstack.com/) for web server routing
- [Zap Logger](https://github.com/uber-go/zap) for logging
- [Sentry](https://sentry.io/) for error tracking and reporting
- [Cron](https://godoc.org/github.com/robfig/cron) for scheduling tasks
- [errors](https://github.com/pkg/errors) for error handling and adding stack trace to golang
- [OZZO](github.com/go-ozzo/ozzo-validation) for data validation

## Getting started

Below we describe the conventions or tools specific to golang project.

### Database

For information about the project's database, see the [DB.md](docs/DB.md) file.

### Protocol Buffer

To use protocol buffer for gRPC communication please refer to [Protohub](https://github.com/infranyx/protobuf-template). Protohub is a hub for managing your protobuf files and with auto generation feature you can simply `go get` the generated code of your proto.

### API docs

The template doesn't have API docs. For auto-generated API docs that you include, you can also give instructions on the
build process.

### Docker Compose

Using Compose is an easy way to manage multi-container applications on any system that supports Docker.

```bash
docker-compose up -d # it runs all necessary docker images that is needed

docker-compose -f docker-compose.e2e-local.yaml up -d # it runs all necessary docker images for Testing environment
```

### Linting

Linting is an important part of any Go project. It helps to ensure that code is written in a consistent and maintainable way, and can help to catch potential errors before they become problems. It is important to note that linting should be done regularly throughout the development process, not just at the end. This will help ensure that any potential issues are caught early on and can be fixed quickly and easily.

To lint your Go project, you can simply use makefile

```makefile
lint-dockerfile: # Lint your Dockerfile

lint-go: # Use golintci-lint on your project
```

## Notes
