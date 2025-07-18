package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var DBConnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
	fmt.Println("1")
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}
	fmt.Println("2")
	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	fmt.Println("3")
	if errs != nil {
		fmt.Println(errs.Error())
		panic(errs)
	}
	fmt.Println("4")
	DBConnection = dbParam

	fmt.Println("Migration success, applied", n, "migrations!")
}
