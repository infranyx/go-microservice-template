package infraStruct

import (
	"github.com/infranyx/go-grpc-template/pkg/config"
	"github.com/infranyx/go-grpc-template/pkg/postgres"
)

type InfraStruct struct {
	Pg  *postgres.Postgres
	Cfg *config.Config
}

var Infra *InfraStruct

func init() {
	Infra = NewInfraStruct()
}

func NewInfraStruct() *InfraStruct {
	return &InfraStruct{Pg: postgres.PgConn, Cfg: config.Conf}
}
