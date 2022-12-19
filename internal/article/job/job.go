package job

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	"github.com/infranyx/go-grpc-template/utils/wrapper"

	cronJob "github.com/infranyx/go-grpc-template/pkg/cron"
)

type articleJob struct {
	cj *cronJob.CronJob
}

func NewArticleJob() articleDomain.ArticleJob {
	cj := cronJob.NewCron()
	return &articleJob{cj: cj}
}

func (aj *articleJob) RunJobs(ctx context.Context) {
	aj.logArticleJob(ctx)
	go aj.cj.Start()
}

func (aj *articleJob) logArticleJob(ctx context.Context) {
	worker := wrapper.BuildChain(aj.logArticleWorker(),
		wrapper.SentryMiddleware,
		wrapper.RecoveryMiddleware,
		wrapper.ErrorHandlerMiddleware,
	)
	entryId, _ := aj.cj.AddFunc("*/1 * * * *",
		worker.ToCronJobFunc(ctx, nil),
	)
	logger.Zap.Sugar().Infof("Starting log article job: %v", entryId)
}
