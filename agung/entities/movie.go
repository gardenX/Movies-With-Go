package entities

import (
	"anggafirdaus10/agung/conf"
	"fmt"
	"time"
)

type Movie struct {
	Id         int
	Title      string
	ReleasedAt string
	Director   Person
	Country    Country
	Genres     []Genre
}

func (m *Movie) Delete() error {
	return conf.Db.QueryRow(fmt.Sprintf(
		"UPDATE movies SET deleted_at = '%s' WHERE movies.id = '%d' AND deleted_at IS NULL RETURNING movies.id",
		time.Now().Format(time.RFC3339),
		m.Id,
	)).Scan(&m.Id)
}

func (m *Movie) LoadDirector() (err error) {
	return conf.Db.QueryRow("SELECT name FROM people WHERE id = $1 AND deleted_at IS NULL", m.Director.Id).Scan(&m.Director.Name)
}

func (m *Movie) LoadGenre() (err error) {
	rows, err := conf.Db.Query("SELECT g.id, g.name FROM genre_movie gm inner join genres g on g.id = gm.genre_id WHERE gm.movie_id = $1 AND g.deleted_at IS NULL", m.Id)

	if err != nil {
		return err
	}

	for rows.Next() {
		genre := Genre{}
		err = rows.Scan(&genre.Id, &genre.Name)

		if err != nil {
			return err
		}

		m.Genres = append(m.Genres, genre)
	}

	return err
}

func (m *Movie) LoadCountry() interface{} {
	return conf.Db.QueryRow("SELECT name FROM countries WHERE id = $1 AND deleted_at IS NULL", m.Country.Id).Scan(&m.Country.Name)
}
