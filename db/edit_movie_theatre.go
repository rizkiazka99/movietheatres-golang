package db

import (
	"movietheatres-go/config"
	"movietheatres-go/models"
)

func UpdateMovieTheatre(id int64, theatre models.MovieTheatre) (int64, error) {
	sqlStatement := `
	UPDATE theatres
	SET nama = $2, lokasi = $3, rating = $4
	WHERE id = $1;`

	res, err := config.Db.Exec(
		sqlStatement,
		id,
		theatre.Nama,
		theatre.Lokasi,
		theatre.Rating,
	)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func DeleteMovieTheatre(id int64) (int64, error) {
	sqlStatement := `DELETE from theatres WHERE id = $1`

	res, err := config.Db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
