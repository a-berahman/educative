package config

import (
	"database/sql"
	"fmt"
	"log"	
	//it supports Postgre configs
	_ "github.com/lib/pq"
)

func GetPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
			CFG.Postgres.Port,
			CFG.Postgres.Host,
			CFG.Postgres.Username,
			CFG.Postgres.Password,
			CFG.Postgres.DBName,
		),
	)
	if err != nil {
		log.Panicln(err)
	}

	db.SetMaxOpenConns(CFG.Postgres.MaxOpenConnections)
	db.SetMaxIdleConns(CFG.Postgres.MaxIdleConnections)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
