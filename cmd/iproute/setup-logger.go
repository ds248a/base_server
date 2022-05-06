package main

import (
	"github.com/pkg/errors"
	// base
	"github.com/ds248a/base/logger"
	//
	"github.com/ds248a/base_server/internal/app"
)

func setupLogger() error {
	ctx := app.Context()
	v, e := app.LoggerLevel.Maybe(ctx)
	if e != nil {
		return errors.Wrap(e, "get logger level from config")
	}
	var l logger.LogLevel
	if e = l.UnmarshalText([]byte(v)); e != nil {
		return errors.Wrapf(e, "recognize '%s' logger level from config", v)
	}
	logger.SetLevel(l)
	return nil
}
