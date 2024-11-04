package middleware

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().Logger()

	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)

		logger.Info().Timestamp().
			Str("method", c.Request().Method).
			Str("path", c.Path()).
			Int("status", c.Response().Status).
			Dur("duration", time.Since(start)).
			Msg("request details")

		return err
	}
}
