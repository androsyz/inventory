package main

import (
	"context"
	"os"

	"github.com/androsyz/inventory/config"
	"github.com/androsyz/inventory/internal/middleware"
	"github.com/androsyz/inventory/internal/server"
	"github.com/rs/zerolog"

	"github.com/androsyz/inventory/internal/pkg/database"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()

	zlog.New(zl)

	cfg := config.NewConfig()
	sqlConn := database.NewDBConnection(*cfg)
	app := server.NewServer(cfg, sqlConn)
	app.ConfigureRoutes()

	app.Echo.Use(middleware.LoggingMiddleware)
	app.Echo.HideBanner = true
	app.Echo.HidePort = true

	zlog.Info(context.Background(), nil, "starting server on :"+cfg.Http.Port)
	err := app.Start(cfg.Http.Port)
	if err != nil {
		zlog.Fatal(context.Background(), nil, "Port already used"+err.Error())
	}
}
