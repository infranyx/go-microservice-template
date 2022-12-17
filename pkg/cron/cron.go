package cronJob

import (
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	*cron.Cron
}

func NewCron() *CronJob {
	cl := NewCronLogger()
	c := cron.New(cron.WithChain(

		cron.SkipIfStillRunning(cl),
	))

	return &CronJob{
		c,
	}
}
