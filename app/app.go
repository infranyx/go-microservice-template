package app

import (
	"context"
	articleConfigurator "github.com/infranyx/go-grpc-template/internal/article/configurator"
	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
	"os"
	"os/signal"
	"syscall"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ic, infraDown, err := iContainer.NewIC(ctx)
	if err != nil {
		return err
	}
	defer infraDown()

	configureModule(ctx, ic)

	var grpcServerError error
	go func() {
		if err := ic.GrpcServer.RunGrpcServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunGrpcServer) err: {%v}", err)
			grpcServerError = err
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		ic.Logger.Sugar().Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		ic.Logger.Sugar().Errorf("ctx.Done: %v", done)
	}

	ic.Logger.Sugar().Info("Server Exited Properly")
	return grpcServerError
}

func configureModule(ctx context.Context, ic *iContainer.IContainer) {
	articleConfigurator.NewArticleConfigurator(ic).ConfigureArticle(ctx)
}
