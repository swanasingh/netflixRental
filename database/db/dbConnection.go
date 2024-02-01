package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"netflixRental/configs"
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
