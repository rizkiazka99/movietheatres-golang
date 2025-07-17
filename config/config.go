package config

import "database/sql"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "postgres"
	Dbname   = "db-movie-theatres-go"
)

var (
	Db  *sql.DB
	Err error
)
