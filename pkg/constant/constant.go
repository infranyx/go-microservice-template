package constant

import "time"

const AppName = "Go-gRPC-Template"

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
)

const (
	HttpPort = 4000
	GrpcPort = 3000
	GrpcHost = "localhost"
)

const (
	PgMaxConn         = 1
	PgMaxIdleConn     = 1
	PgMaxLifeTimeConn = 1
	PgSslMode         = "disable"
)

const (
	MaxHeaderBytes       = 1 << 20
	BodyLimit            = "2M"
	ReadTimeout          = 15 * time.Second
	WriteTimeout         = 15 * time.Second
	GzipLevel            = 5
)
