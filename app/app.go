package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	article_configurator "github.com/infranyx/go-grpc-template/internal/article/configurator"
	iContainer "github.com/infranyx/go-grpc-template/shared/infra_container"
)

type App struct {
}

// Server constructor
func New() *App {
	return &App{}
}

func (a *App) Run() error {
	var serverError error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ic, infraCleanup, err := iContainer.NewIC(ctx)
	if err != nil {
		return err
	}
	defer infraCleanup()

	//
	articleConfigurator := article_configurator.NewArticleConfigurator(ic)
	articleConfigurator.ConfigureArticle(ctx)

	//

	go func() {
		if err := ic.GrpcServer.RunGrpcServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunGrpcServer) err: {%v}", err)
			serverError = err
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
	return serverError
}
