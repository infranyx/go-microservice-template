package tmpDirHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
)

type tmpDirHealthCheck struct {
}

func NewTmpDirHealthCheck() healthCheckDomain.TmpDirHealthCheckUseCase {
	return &tmpDirHealthCheck{}
}

func (th *tmpDirHealthCheck) PingCheck() bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	tmpDirPath := filepath.Join(cwd, "/tmp")
	if unix.Access(tmpDirPath, unix.W_OK) != nil {
		return false
	}

	return true
}
