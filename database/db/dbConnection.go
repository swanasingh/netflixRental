package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"
	"log"
	"netflixRental/configs"
	"os"
)

func CreateConnection(cfg configs.Config) *sql.DB {
	configs.GetConfigs(&cfg)
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
		cfg.Database.User,
		cfg.Database.Password,
	)
	fmt.Println(dataSourceName)
	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}

func RunMigration(cfg configs.Config) {
	// Initialize migration instance
	source, _ := os.Getwd()
	path := fmt.Sprintf("%s/database/migration", source)
	fmt.Println(path)
	m, err := migrate.New(
		path,
		"postgres://root:pass@localhost:5432/netflix-rental?sslmode=disable",
	)
	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatalf("Error initializing migration: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}

//$ migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
