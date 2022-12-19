package job

import (
	"context"

	articleDomain "github.com/infranyx/go-grpc-template/internal/article/domain"
	errorUtils "github.com/infranyx/go-grpc-template/pkg/error/error_utils"
	"github.com/infranyx/go-grpc-template/pkg/logger"

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
	worker := errorUtils.HandlerErrorWrapper(
		ctx,
		aj.logArticleWorker(ctx),
	)
	entryId, _ := aj.cj.AddFunc("*/1 * * * *",
		worker,
	)
	logger.Zap.Sugar().Infof("Starting log article job: %v", entryId)
}
