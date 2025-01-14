package src_test

import (
	"context"
	"errors"
	"euro-exchange/config"
	"euro-exchange/src"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestApp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := NewMockClient(ctrl)

	logger := logrus.StandardLogger()
	hook := new(Hook)
	logger.AddHook(hook)
	testLogger := &src.Logger{
		StdLogger:  logger,
		FileLogger: logger,
	}

	app := src.NewApp(&config.Config{ChecksFrequency: 30}, testLogger)
	app.Client = client

	t.Run("success run", func(t *testing.T) {
		defer func() { hook.Entries = []logrus.Entry{} }()
		client.EXPECT().GetExchangeRate().Return(&src.ExchangeRates{
			Table:    "a",
			Currency: "euro",
			Code:     "EUR",
			Rates: []src.Rate{
				{
					No:   "no-1",
					Date: "2024-01-01",
					Mid:  4.45,
				},
				{
					No:   "no-2",
					Date: "2024-01-02",
					Mid:  4.60,
				},
			},
		}, nil)

		wg := &sync.WaitGroup{}
		wg.Add(1)
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		app.Run(ctx, wg)

		assert.Equal(t, 2, len(hook.Entries))
		assert.Contains(t, hook.Entries[0].Message, "graceful runner shutdown")
	})

	t.Run("error getting exchange rates", func(t *testing.T) {
		defer func() { hook.Entries = []logrus.Entry{} }()
		client.EXPECT().GetExchangeRate().Return(nil, errors.New("test error"))

		wg := &sync.WaitGroup{}
		wg.Add(1)
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		app.Run(ctx, wg)

		assert.Equal(t, 4, len(hook.Entries))
		assert.Contains(t, hook.Entries[0].Message, "test error")
	})
}

type Hook struct {
	Entries []logrus.Entry
	mu      sync.RWMutex
}

func (t *Hook) Fire(e *logrus.Entry) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Entries = append(t.Entries, *e)
	return nil
}

func (t *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
