# Go Microservice Template

<!-- # Hands-On Microservices with Node.js -->

This is the code repository for [Go Microservice Template](https://www.github.com/infranyx/go-microservice-template), published by InfraNyx.

**Build, test, and deploy robust golang microservices**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About the project](#about-the-project)
- [Status](#status)
- [Reference](#reference)
- [Getting started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
- [Testing](#testing)
  - [Linting](#linting)
  - [Building](#building)
  - [Cleaning](#cleaning)
  - [Continuous Integration](#continuous-integration)
- [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## About the project

The template is used to create golang project. All golang projects must follow the conventions in the
template. Calling for exceptions must be brought up in the engineering team.

## Status

The template project is in alpha status.

## Reference

- [Database Schema and Migrations](docs/db.md)
- [Deployment Instructions](docs/deployment.md)
- [Design Decisions and Technical Considerations](docs/design.md)
- [Development and Operations Processes](docs/devops.md)
- [Environment Variables](docs/env.md)

## Getting started

### Prerequisites

- [Go 1.13+](https://golang.org/doc/install)

### Setup

1. Clone the repository and navigate to the root directory of the project:

```bash
git clone https://github.com/infranyx/go-microservice-template.git
cd go-microservice-template
```

2. Install the dependencies:

```bash
make dep
```

3. Run the development server:

```bash
make run_dev
```

This will start the http server on port 4000 and grpc server on port 3000. You can customize the ports by setting the `HTTP_PORT` and `GRPC_PORT` environment variable.

## Testing

To run the tests, use the following command:

```bash
make test
```

To generate a test coverage report, use the following command:

```bash
make test_coverage
```

### Linting

To lint the code, use the following command:

```bash
make lint
```

This will run all available linters, including Go lint, Dockerfile lint, and YAML lint.

### Building

To build the binary, use the following command:

```bash
make build
```

This will create a binary in the `out/bin` directory. You can run the binary with the following command:

```bash
make run
```

### Cleaning

To clean the build artifacts and generated files, use the following command:

```bash
make clean
```

This will remove the `bin` and `out` directories, as well as any build-related files.

### Continuous Integration

To run the continuous integration tasks, use the following command:

```bash
make ci
```

This will run the tests, linting, and code coverage tasks, and generate the corresponding reports. The reports will be saved to the `out` directory.

## Notes
