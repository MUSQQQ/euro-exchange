package main

import (
	"context"
	"euro-exchange/config"
	"euro-exchange/src"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()

	logger := src.NewLogger(cfg)

	logger.Log(logrus.InfoLevel, logrus.Fields{
		"app_name": "euro-exchange",
		"message":  "app start",
	})

	app := src.NewApp(cfg, logger)

	ticker := time.Tick(time.Duration(cfg.Timeout) * time.Second)

	ctx := context.Background()
	cancCtx, cancFunc := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}

	for i := 0; i < cfg.ChecksNumber; i++ {
		wg.Add(1)
		go app.Run(cancCtx, wg)
	}

	<-ticker
	cancFunc()
	wg.Wait()

	logger.Log(logrus.InfoLevel, logrus.Fields{
		"message": "graceful app shutdown",
	})
}
