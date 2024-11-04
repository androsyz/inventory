package server

import (
	"database/sql"

	"github.com/androsyz/inventory/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Sql    *sql.DB
	Config *config.Config
}

func NewServer(cfg *config.Config, sql *sql.DB) *Server {
	return &Server{
		Echo:   echo.New(),
		Sql:    sql,
		Config: cfg,
	}
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
