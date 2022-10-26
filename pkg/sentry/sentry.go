package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
)

type Sentry struct {
	Client *sentry.Client
}

func NewSentryClient() *Sentry {
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              "https://60c0bc85835b45adadb29d6ad91ac861@sentry.zarinworld.ir/18",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		panic(err)
	}
	defer client.Flush(2 * time.Second)
	return &Sentry{
		Client: client,
	}
}

func (s *Sentry) CaptureException(exception error) {
	tags := make(map[string]string)
	localHub := sentry.CurrentHub().Clone()
	localHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(tags)
	})
	localHub.CaptureException(exception)
}
