package tmpDirHealthCheckUseCase

import (
	healthCheckDomain "github.com/infranyx/go-grpc-template/internal/health_check/domain"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"golang.org/x/sys/unix"
	"path/filepath"
	"runtime"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.TmpDirHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
	if !config.IsProdEnv() {
		return true
	}

	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		return false
	}

	tmpDir := filepath.Join(filepath.Dir(callerDir), "../../../..", "tmp")

	if unix.Access(tmpDir, unix.W_OK) != nil {
		return false
	}

	return true
}
