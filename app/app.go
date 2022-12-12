package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	articleConfigurator "github.com/infranyx/go-grpc-template/internal/article/configurator"
	cContainer "github.com/infranyx/go-grpc-template/pkg/client_container"
	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
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

	cc, clientsDown, err := cContainer.NewCC(ctx)
	if err != nil {
		return err
	}
	defer clientsDown()

	me := configureModule(ctx, ic, cc)
	if me != nil {
		return me
	}

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

func configureModule(ctx context.Context, ic *iContainer.IContainer, cc *cContainer.CContainer) error {
	e := articleConfigurator.NewArticleConfigurator(ic, cc).ConfigureArticle(ctx)
	if e != nil {
		return e
	}
	return nil
}
