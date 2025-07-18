package db

import (
	"database/sql"
	"fmt"
	"movietheatres-go/config"
	"movietheatres-go/models"
)

func GetMovieTheatres() ([]models.MovieTheatre, error) {
	var results []models.MovieTheatre

	sqlStatement := `SELECT * from theatres`

	rows, err := config.Db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var theatre = models.MovieTheatre{}

		err = rows.Scan(
			&theatre.ID,
			&theatre.Nama,
			&theatre.Lokasi,
			&theatre.Rating,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, theatre)
	}

	fmt.Println("Data:", results)
	return results, nil
}

func GetMovieTheatreById(id int64) (models.MovieTheatre, error) {
	var theatre models.MovieTheatre

	sqlStatement := `SELECT * from theatres WHERE id = $1`

	row := config.Db.QueryRow(sqlStatement, id)

	err := row.Scan(
		&theatre.ID,
		&theatre.Nama,
		&theatre.Lokasi,
		&theatre.Rating,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return theatre, nil
		} else {
			return theatre, err
		}
	}

	return theatre, nil
}
