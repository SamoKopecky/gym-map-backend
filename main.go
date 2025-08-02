package main

import (
	"flag"
	"gym-map/api/app"
	"gym-map/config"
	"gym-map/db"
	"log"

	"github.com/joho/godotenv"
)

const MIGRATION_PATH = "file://migrations"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v. Continuing without .env", err)
	}
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
