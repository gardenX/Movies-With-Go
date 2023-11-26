package repositories

import (
	"anggafirdaus10/agung/conf"
	"anggafirdaus10/agung/entities"
	"fmt"
	"strings"
)

func (r movieRepository) Search(request MovieSearchRequest) (movies []entities.Movie, err error) {
	q := []string{
		"SELECT m.id, m.title, m.released_at, p.name, c.name, g.name FROM movies m INNER JOIN people p on p.id = m.director_id INNER JOIN countries c on m.country_id = c.id inner join genre_movie gm on m.id = gm.movie_id inner join genres g on g.id = gm.genre_id WHERE m.deleted_at IS NULL",
	}

	if request.Title != "" {
		q = append(q, fmt.Sprintf("m.title ILIKE '%[1]s%[2]s%[1]s'", "%", request.Title))
	}

	if request.Year != "" {
		q = append(q, fmt.Sprintf("m.released_at BETWEEN '%[1]s-01-01' AND '%[1]s-12-31'", request.Year))
	}

	if request.Genre != "" {
		q = append(q, fmt.Sprintf("EXISTS (SELECT g.* FROM genres g INNER JOIN genre_movie gm ON g.id = gm.genre_id WHERE m.id = gm.movie_id AND g.deleted_at IS NULL AND g.name ilike '%[1]s%[2]s%[1]s')", "%", request.Genre))
	}

	query := strings.Join(q, " AND ")

	rows, err := conf.Db.Query(query)
	keys := map[int]int{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		movie := entities.Movie{}
		genre := entities.Genre{}

		err = rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.ReleasedAt,
			&movie.Director.Name,
			&movie.Country.Name,
			&genre.Name,
		)

		if err != nil {
			return nil, err
		}

		_, ok := keys[movie.Id]

		if !ok {
			keys[movie.Id] = len(keys)
			movie.Genres = []entities.Genre{genre}
			movies = append(movies, movie)
		} else {
			movies[keys[movie.Id]].Genres = append(movies[keys[movie.Id]].Genres, genre)
		}
	}

	return
}

type MovieSearchRequest struct {
	Title string
	Year  string
	Genre string
}
