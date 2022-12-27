package articleJob

import (
	"context"
	"go.uber.org/zap"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
	wrapperErrorhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/error_handler"
	wrapperRecoveryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/recovery_handler"
	wrapperSentryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/sentry_handler"
	"github.com/robfig/cron/v3"

	cronJob "github.com/infranyx/go-grpc-template/pkg/cron"
)

type job struct {
	cron   *cron.Cron
	logger *zap.Logger
}

func NewJob(logger *zap.Logger) articleDomain.ArticleJob {
	cron := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cronJob.NewCronLogger()),
	))
	return &job{cron: cron, logger: logger}
}

func (j *job) StartJobs(ctx context.Context) {
	j.logArticleJob(ctx)
	go j.cron.Start()
}

func (j *job) logArticleJob(ctx context.Context) {
	worker := wrapper.BuildChain(j.logArticleWorker(),
		wrapperSentryhandler.SentryHandler,
		wrapperRecoveryhandler.RecoveryHandler,
		wrapperErrorhandler.ErrorHandler,
	)
	entryId, _ := j.cron.AddFunc("*/1 * * * *",
		worker.ToCronJobFunc(ctx, nil),
	)
	j.logger.Sugar().Infof("Starting log article job: %v", entryId)
}
