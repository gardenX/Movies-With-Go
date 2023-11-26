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

func (m *Movie) Find() (err error) {
	return conf.Db.
		QueryRow("SELECT m.id, m.title, m.released_at, m.director_id, m.country_id FROM movies m WHERE m.id = $1 AND deleted_at is NULL", m.Id).
		Scan(&m.Id, &m.Title, &m.ReleasedAt, &m.Director.Id, &m.Country.Id)
}

func (m *Movie) Create() (err error) {
	err = conf.Db.QueryRow(
		"insert into movies (title, released_at, created_at, updated_at, director_id, country_id) values ($1, $2, $3, $4, $5, $6) RETURNING id, released_at",
		m.Title,
		m.ReleasedAt,
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		m.Director.Id,
		m.Country.Id,
	).Scan(&m.Id, &m.ReleasedAt)

	if err != nil {
		return
	}

	err = m.attachGenres()

	return
}

func (m *Movie) attachGenres() (err error) {
	for _, genre := range m.Genres {
		err = conf.Db.QueryRow(
			"INSERT INTO genre_movie (genre_id, movie_id) VALUES ($1, $2) RETURNING genre_id",
			genre.Id,
			m.Id,
		).Scan(&genre.Id)

		if err != nil {
			return
		}
	}

	return err
}

func (m *Movie) Update() (err error) {
	err = conf.Db.QueryRow(
		"UPDATE movies SET title = $1, released_at = $2, updated_at = $3, director_id = $4, country_id = $5 WHERE id = $6 AND deleted_at IS NULL RETURNING released_at",
		m.Title,
		m.ReleasedAt,
		time.Now().Format(time.RFC3339),
		m.Director.Id,
		m.Country.Id,
		m.Id,
	).Scan(&m.ReleasedAt)

	if err != nil {
		return
	}

	conf.Db.QueryRow("DELETE FROM genre_movie WHERE movie_id = $1", m.Id)
	err = m.attachGenres()

	return
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
