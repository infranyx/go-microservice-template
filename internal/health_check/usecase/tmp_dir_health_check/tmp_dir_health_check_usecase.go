package tmpDirHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.TmpDirHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
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
