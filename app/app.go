package app

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/infranyx/go-grpc-template/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	articleConfigurator "github.com/infranyx/go-grpc-template/internal/article/configurator"
	cContainer "github.com/infranyx/go-grpc-template/pkg/client_container"
	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	se := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Conf.Sentry.Dsn,
		TracesSampleRate: 1.0,
	})
	if se != nil {
		log.Fatalf("sentry.Init: %s", se)
	}
	defer sentry.Flush(2 * time.Second)

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

	var serverError error
	go func() {
		if err := ic.GrpcServer.RunGrpcServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunGrpcServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	go func() {
		if err := ic.EchoServer.RunHttpServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunEchoServer) err: {%v}", err)
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

func configureModule(ctx context.Context, ic *iContainer.IContainer, cc *cContainer.CContainer) error {
	e := articleConfigurator.NewArticleConfigurator(ic, cc).ConfigureArticle(ctx)
	if e != nil {
		return e
	}
	return nil
}
