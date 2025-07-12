package main

import (
	"flag"
	"gym-map/api/app"
	"gym-map/config"
	"gym-map/db"
)

const MIGRATION_PATH = "file://migrations"

func main() {
	cfg := config.GetConfig()
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	dbConn := db.GetDbConn(cfg.GetDSN(), *debug, MIGRATION_PATH)
	dbConn.RunMigrations()

	go func() {
		app.RunMetricsApi()
	}()
	app.RunApi(dbConn.Conn, &cfg)
}
