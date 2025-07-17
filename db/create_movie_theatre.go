package db

import (
	"fmt"
	"movietheatres-go/config"
	"movietheatres-go/models"
)

func CreateMovieTheatre(theatre models.MovieTheatre) {
	var movieTheatre models.MovieTheatre

	sqlStatement := `
	INSERT INTO theatres (nama, lokasi, rating)
	VALUES ($1, $2, $3)
	Returning *
	`

	config.Err = config.Db.QueryRow(
		sqlStatement,
		theatre.Nama,
		theatre.Lokasi,
		theatre.Rating,
	).Scan(&movieTheatre.ID, &movieTheatre.Nama, &movieTheatre.Lokasi, &movieTheatre.Rating)

	if config.Err != nil {
		panic(config.Err)
	} else {
		fmt.Printf("New Movie Theatre Data: %+v\n", movieTheatre)
	}
}
