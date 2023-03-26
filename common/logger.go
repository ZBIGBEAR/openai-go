package common

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var log = zerolog.New(os.Stderr).With().Logger()

type Logger interface {
	Infof(ctx context.Context, format string, args ...any)
	Warnf(ctx context.Context, format string, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
}

type logger struct{}

var _ Logger = &logger{}

func NewLog() Logger {
	return &logger{}
}

func (l *logger) baseLog(ctx context.Context, log *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		ctx = context.TODO()
	}
	logCtx := log.Time(LogTime, time.Now())
	return logCtx.Interface(ServiceName, ctx.Value(ServiceName))
}

func (l *logger) Infof(ctx context.Context, format string, args ...any) {
	l.baseLog(ctx, log.Info()).Str(LogMsg, fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(ctx context.Context, format string, args ...any) {
	l.baseLog(ctx, log.Warn()).Str(LogMsg, fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(ctx context.Context, format string, args ...any) {
	l.baseLog(ctx, log.Error()).Str(LogMsg, fmt.Sprintf(format, args...))
}
