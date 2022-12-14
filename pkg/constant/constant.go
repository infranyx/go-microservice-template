package constant

import "time"

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
	StackSize            = 1 << 10 // 1 KB
	BodyLimit            = "2M"
	ReadTimeout          = 15 * time.Second
	WriteTimeout         = 15 * time.Second
	GzipLevel            = 5
	WaitShotDownDuration = 3 * time.Second
	Dev                  = "development"
	Test                 = "test"
	Production           = "production"
)
