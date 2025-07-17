package main

import (
	"database/sql"
	"fmt"
	"movietheatres-go/config"
	"movietheatres-go/router"

	_ "github.com/lib/pq"
)

func main() {
	startServer()
	// theatre := models.MovieTheatre{
	// 	ID:     utils.IDGenerator(),
	// 	Nama:   "Metropole XXI",
	// 	Lokasi: "Megaria, Komplek, Jl. Pegangsaan Barat No.21, RT.1/RW.1, Pegangsaan, Kec. Menteng, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10320",
	// 	Rating: 4.2,
	// }
	// defer db.CreateMovieTheatre(theatre)
	// connectToDB()
}

func connectToDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname,
	)
	config.Db, config.Err = sql.Open("postgres", psqlInfo)
	if config.Err != nil {
		panic(config.Err)
	}
	// defer config.Db.Close()

	config.Err = config.Db.Ping()
	if config.Err != nil {
		panic(config.Err)
	}

	fmt.Println("Successfully connected to the database")
}

func startServer() {
	var PORT = ":8080"

	connectToDB()
	router.StartServer().Run(PORT)
}
