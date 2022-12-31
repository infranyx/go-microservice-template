# Environment Variables

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

This file lists the environment variables that are used by the project.

| Variable                         | Description                                                             | Required | Default Value    |
| -------------------------------- | ----------------------------------------------------------------------- | -------- | ---------------- |
| `APP_ENV`                        | The environment that the application is running in                      | Yes      | `prod`           |
| `GRPC_HOST`                      | The hostname for the gRPC server                                        | Yes      | `localhost`      |
| `GRPC_PORT`                      | The port for the gRPC server                                            | Yes      | `3000`           |
| `HTTP_PORT`                      | The port for the HTTP server                                            | Yes      | N/A              |
| `EXTERNAL_GO_TEMPLATE_GRPC_PORT` | The port for an external gRPC server                                    | Yes      | `3000`           |
| `EXTERNAL_GO_TEMPLATE_GRPC_HOST` | The hostname for an external gRPC server                                | Yes      | `localhost`      |
| `PG_HOST`                        | The hostname for the Postgresql server                                  | Yes      | `localhost`      |
| `PG_PORT`                        | The port for the Postgresql server                                      | Yes      | `5432`           |
| `PG_USER`                        | The username for the Postgresql server                                  | Yes      | `admin`          |
| `PG_PASS`                        | The password for the Postgresql server                                  | Yes      | `admin`          |
| `PG_DB`                          | The database name for the Postgresql server                             | Yes      | `grpc_template`  |
| `PG_MAX_CONNECTIONS`             | The maximum number of connections allowed to the Postgresql server      | Yes      | `1`              |
| `PG_MAX_IDLE_CONNECTIONS`        | The maximum number of idle connections allowed to the Postgresql server | Yes      | `1`              |
| `PG_MAX_LIFETIME_CONNECTIONS`    | The maximum lifetime of connections to the Postgresql server            | Yes      | `1`              |
| `PG_SSL_MODE`                    | The SSL mode for the Postgresql server                                  | Yes      | `disable`        |
| `KAFKA_ENABLED`                  | Whether or not Kafka is enabled                                         | Yes      | `1`              |
| `KAFKA_LOG_EVENTS`               | Whether or not to log events to Kafka                                   | Yes      | `1`              |
| `KAFKA_CLIENT_ID`                | The client ID for Kafka                                                 | Yes      | `dev-consumer`   |
| `KAFKA_CLIENT_GROUP_ID`          | The client group ID for Kafka                                           | Yes      | `dev-group`      |
| `KAFKA_CLIENT_BROKERS`           | The brokers for Kafka                                                   | Yes      | `localhost:9094` |
| `KAFKA_NAMESPACE`                | The namespace for Kafka                                                 | Yes      | `dev`            |
| `KAFKA_TOPIC`                    | The topic for Kafka                                                     | Yes      | `test-topic`     |
| `SENTRY_DSN`                     | The Data Source Name (DSN) for Sentry                                   | Yes      | `*`              |
