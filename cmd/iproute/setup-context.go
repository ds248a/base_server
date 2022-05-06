package main

import (
	"context"
	"go.uber.org/zap"
	//base
	"github.com/ds248a/base/logger"
	"github.com/ds248a/base/pkg/signals"
	//
	"github.com/ds248a/base_server/internal/app"
)

func setupContext() {
	ctx, cancel := context.WithCancel(context.Background())
	signals.WhenSignalExit(func() error {
		logger.SetLevel(zap.InfoLevel)
		logger.Info(ctx, "caught application stop signal")
		cancel()
		return nil
	})
	app.SetContext(ctx)
}
