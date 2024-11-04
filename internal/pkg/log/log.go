package log

import (
	"context"

	"github.com/rs/zerolog"
)

var log *Zlog

type ZlogLevel int

const (
	InfoLevel  ZlogLevel = ZlogLevel(zerolog.InfoLevel)
	ErrorLevel ZlogLevel = ZlogLevel(zerolog.ErrorLevel)
	FatalLevel ZlogLevel = ZlogLevel(zerolog.FatalLevel)
	DebugLevel ZlogLevel = ZlogLevel(zerolog.DebugLevel)
	WarnLevel  ZlogLevel = ZlogLevel(zerolog.WarnLevel)
	PanicLevel ZlogLevel = ZlogLevel(zerolog.PanicLevel)
	TraceLevel ZlogLevel = ZlogLevel(zerolog.TraceLevel)
)

type Zlog struct {
	log zerolog.Logger
}

type ZlogProperties struct {
	Level      ZlogLevel
	AppName    string
	AppVersion string
	AppEnv     string
}

func New(logz zerolog.Logger) *Zlog {
	log = &Zlog{
		log: logz,
	}

	return log
}

func Info(ctx context.Context, data interface{}, msg string) {
	log.log.Info().Fields(data).Msg(msg)
}

func Error(ctx context.Context, data interface{}, msg string) {
	log.log.Error().Fields(data).Msg(msg)
}

func Fatal(ctx context.Context, data interface{}, msg string) {
	log.log.Fatal().Fields(data).Msg(msg)
}

func Debug(ctx context.Context, data interface{}, msg string) {
	log.log.Debug().Fields(data).Msg(msg)
}

func Warn(ctx context.Context, data interface{}, msg string) {
	log.log.Warn().Fields(data).Msg(msg)
}

func Panic(ctx context.Context, data interface{}, msg string) {
	log.log.Panic().Fields(data).Msg(msg)
}

func Trace(ctx context.Context, data interface{}, msg string) {
	log.log.Trace().Fields(data).Msg(msg)
}
