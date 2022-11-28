package constant

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
)

const (
	GrpcPort = 3000
	GrpcHost = "localhost"
)

const (
	PgHost            = "localhost"
	PgPort            = 5432
	PgUser            = "admin"
	PgPass            = "admin"
	PgDb              = "go-grpc-template"
	PgMaxConn         = 1
	PgMaxIdleConn     = 1
	PgMaxLifeTimeConn = 1
	PgSslMode         = "disable"
)
