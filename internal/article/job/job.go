package articleJob

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/pkg/wrapper"
	wrapperErrorhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/error_handler"
	wrapperRecoveryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/recovery_handler"
	wrapperSentryhandler "github.com/infranyx/go-grpc-template/pkg/wrapper/handlers/sentry_handler"
	"github.com/robfig/cron/v3"

	cronJob "github.com/infranyx/go-grpc-template/pkg/cron"
)

type articleJob struct {
	cron *cron.Cron
}

func NewArticleJob() articleDomain.ArticleJob {
	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cronJob.NewCronLogger()),
	))
	return &articleJob{cron: c}
}

func (j *articleJob) StartJobs(ctx context.Context) {
	j.logArticleJob(ctx)
	go j.cron.Start()
}

func (j *articleJob) logArticleJob(ctx context.Context) {
	worker := wrapper.BuildChain(j.logArticleWorker(),
		wrapperSentryhandler.SentryHandler,
		wrapperRecoveryhandler.RecoveryHandler,
		wrapperErrorhandler.ErrorHandler,
	)
	entryId, _ := j.cron.AddFunc("*/1 * * * *",
		worker.ToCronJobFunc(ctx, nil),
	)
	logger.Zap.Sugar().Infof("Starting log article job: %v", entryId)
}
