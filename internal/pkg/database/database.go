package database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	config "github.com/androsyz/inventory/config"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
	_ "github.com/lib/pq"
)

func NewDBConnection(cfg config.Config) *sql.DB {
	zlog.Info(context.Background(), nil, "init postgresql instance")

	dbHost := cfg.Database.Host
	port := cfg.Database.Port
	dbUser := cfg.Database.User
	dbPassword := cfg.Database.Password
	dbName := cfg.Database.Name

	dbPort, _ := strconv.Atoi(port)

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		zlog.Panic(context.Background(), nil, "failed connect to postgresql: "+err.Error())
		panic(err)
	}

	zlog.Info(context.Background(), nil, "success connect to postgresql")

	return db

}
