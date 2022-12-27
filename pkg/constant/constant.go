package constant

const AppName = "Go-gRPC-Template"

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
)

const (
	HttpHost = "localhost"
	HttpPort = 4000

	GrpcHost = "localhost"
	GrpcPort = 3000
)

const (
	PgMaxConn         = 1
	PgMaxIdleConn     = 1
	PgMaxLifeTimeConn = 1
	PgSslMode         = "disable"
)

const (
	GzipLevel = 5
)
