# Database

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Migrations](#migrations)
  - [Installation](#installation)
  - [Usage](#usage)
- [PostgreSQL](#postgresql)
  - [Installation](#installation-1)
  - [Connecting to the Database](#connecting-to-the-database)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

This project uses PostgreSQL as its primary data store, and Migrate for managing database migrations.

## Migrations

This code repository use [Migrate](https://github.com/golang-migrate/migrate) as a database migration tool. It helps you manage and keep track of database schema changes in your Go applications.

### Installation

To install migrate cli please go to [installation guide](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

### Usage

You can use the following command to set the PG_URL environment variable and then use it in the migrate command:

```bash
export PG_URL=postgres://[user]:[password]@[host]:[port]/[database]
```

This can be useful if you need to run the command multiple times and don't want to type out the full database URL each time.

To create a new migration, use the following command:

```makefile
migrate-create: # Target for running 'create'
	migrate create -ext sql -dir db/migrations $(NAME)
```

This will create a new migration file in the db/migrations directory with the name [timestamp]_[migration_name].up.sql and [timestamp]_[migration_name].down.sql, which contain the SQL statements for applying and rolling back the migration, respectively.

To run the migrations, use the following command:

```makefile
migrate-up:   # Target for running 'up' command
 migrate -path db/migrations -database  $(PG_URL) up
```

To rollback the last migration, use the following command:

```makefile
migrate-down: # Target for running 'down' command
	migrate -path db/migrations -database $(PG_URL) down
```

To drop all tables and sequences in the database, use the following command:

```makefile
migrate-drop: # Target for running 'drop' command
	migrate drop -database $(PG_URL)
```

**Note**: This command will permanently delete all data in the database, so use caution when running it.

To apply or rollback a migration to a specific version , use the following command:

```makefile
migrate-force: # Target for running 'force' command
 migrate -path db/migrations -database $(PG_URL) force $(VERSION)
```

This is useful when you want to undo or redo a specific migration, or when you want to apply a migration that was previously rolled back.
**Note**: The migrate force command should be used with caution, as it can permanently alter the state of the database. Make sure you have a backup of your data before using this command.

## PostgreSQL

PostgreSQL is a powerful, open-source object-relational database system with a strong reputation for reliability, feature robustness, and performance. It is commonly used as the primary data store for web, mobile, geospatial, and analytics applications.

Some key features of PostgreSQL include:

- Support for multiple data types, including text, numerical, and spatial data
- Support for ACID transactions, which ensure the consistency and integrity of data
- Support for triggers, stored procedures, and views, which allow you to define custom logic and data manipulation operations
- Support for full-text search and advanced indexing options, which enable fast querying and data retrieval
- Support for JSON and JSONB data types, which allow you to store and manipulate complex, nested data structures

### Installation

To get started with PostgreSQL, you will need to install the database server and client libraries on your machine. There are various ways to install PostgreSQL, including using a package manager, downloading the binaries from the official website or using an existing Docker Compose file.

Follow these steps:

1. Make sure that Docker and Docker Compose are installed on your machine.
2. Open the docker-compose.yml file in your project directory.
3. If the PostgreSQL service definition already exists in Docker Compose file, you can start the PostgreSQL container by running the following command:

```bash
docker-compose up -d
```

This will start the PostgreSQL container in detached mode.

To stop the container, use the following command:

```bash
docker-compose stop
```

### Connecting to the Database

Replace myuser, mypassword, and mydatabase with the desired username, password, and database name in the docker-coompose file.

To connect to the PostgreSQL database from your application, use the following connection string:

```
postgres://myuser:mypassword
```

Once you have installed PostgreSQL, you can create a database and start using it in your application by connecting to it using a database driver, such as the pgx driver for Go.

To learn more about PostgreSQL, you can refer to the official documentation at https://www.postgresql.org/docs/.
