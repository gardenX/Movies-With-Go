package repositories

import (
	"anggafirdaus10/agung/conf"
	"anggafirdaus10/agung/entities"
	"fmt"
)

type GenreSearchRequest struct {
	Name string
}

func (g genreRepository) Find(id string) (genre entities.Genre, err error) {
	err = conf.Db.
		QueryRow("SELECT id, name FROM genres WHERE id = $1 AND deleted_at is NULL", id).
		Scan(&genre.Id, &genre.Name)

	return
}

func (g genreRepository) Search(request GenreSearchRequest) (genres []entities.Genre, err error) {
	genres = []entities.Genre{}
	rows, err := conf.Db.Query(fmt.Sprintf(
		"SELECT g.id, g.name FROM genres g WHERE g.deleted_at IS NULL and g.name ILIKE '%[1]s%[2]s%[1]s'",
		"%",
		request.Name,
	))

	if err != nil {
		return []entities.Genre{}, err
	}

	for rows.Next() {
		genre := entities.Genre{}
		err = rows.Scan(&genre.Id, &genre.Name)

		if err != nil {
			return []entities.Genre{}, err
		}

		genres = append(genres, genre)
	}

	return genres, err
}
