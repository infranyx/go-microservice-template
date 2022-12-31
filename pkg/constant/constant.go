package constant

// App
const AppName = "Go-Microservice-Template"

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
)

// Http + Grpc
const (
	HttpHost      = "localhost"
	HttpPort      = 4000
	EchoGzipLevel = 5

	GrpcHost = "localhost"
	GrpcPort = 3000
)

// Postgres
const (
	PgMaxConn         = 1
	PgMaxIdleConn     = 1
	PgMaxLifeTimeConn = 1
	PgSslMode         = "disable"
)
