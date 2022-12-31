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

func RecoverWithSentry(hub *sentry.Hub, ctx context.Context, options *Options) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(ctx, err)
		if eventID != nil && options.WaitForDelivery {
			hub.Flush(options.Timeout)
		}

		if options.Repanic {
			panic(err)
		}
	}
}
