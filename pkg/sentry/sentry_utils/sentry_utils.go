package sentryUtils

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
)

type Options struct {
	Repanic         bool
	WaitForDelivery bool
	Timeout         time.Duration
}

func RecoverWithSentry(hub *sentry.Hub, ctx context.Context, o *Options) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(ctx, err)
		if eventID != nil && o.WaitForDelivery {
			hub.Flush(o.Timeout)
		}

		if o.Repanic {
			panic(err)
		}
	}
}
