package commands

import (
	"anggafirdaus10/agung/entities"
	"anggafirdaus10/agung/helpers"
	"anggafirdaus10/agung/repositories"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	movieFlags = map[string]cli.Flag{
		"id":        &cli.IntFlag{Name: "id"},
		"title":     &cli.StringFlag{Name: "title", Aliases: []string{"t"}},
		"released":  &cli.StringFlag{Name: "released", Aliases: []string{"r"}},
		"year":      &cli.StringFlag{Name: "year", Aliases: []string{"y"}},
		"genre":     &cli.StringFlag{Name: "genre", Aliases: []string{"g"}},
		"genreIds":  &cli.IntSliceFlag{Name: "genreIds", Aliases: []string{"gi"}},
		"countryId": &cli.IntFlag{Name: "countryId", Aliases: []string{"ci"}},
		"personId":  &cli.IntFlag{Name: "personId", Aliases: []string{"pi"}},
	}
	MovieCommands = []*cli.Command{
		movieFindCommand(),
		movieSearchCommand(),
		movieCreateCommand(),
		movieUpdateCommand(),
		movieDeleteCommand(),
	}
)

func movieFindCommand() *cli.Command {
	return &cli.Command{
		Name:  "find",
		Usage: "find a movies by id",
		Flags: []cli.Flag{
			movieFlags["id"],
		},
		Action: func(cCtx *cli.Context) error {
			movie := entities.Movie{Id: cCtx.Int("id")}
			err := movie.Find()
			var movies []entities.Movie

			if err == nil {
				_ = movie.LoadDirector()
				_ = movie.LoadGenre()
				_ = movie.LoadCountry()
				movies = append(movies, movie)
			}

			helpers.PrintMoviesTable(movies)

			return nil
		},
	}
}

func movieSearchCommand() *cli.Command {
	return &cli.Command{
		Name:  "search",
		Usage: "search movies",
		Flags: []cli.Flag{
			movieFlags["title"],
			movieFlags["year"],
			movieFlags["genre"],
		},
		Action: func(cCtx *cli.Context) error {
			movies, err := repositories.MovieRepository.Search(repositories.MovieSearchRequest{
				Title: cCtx.String("title"),
				Year:  cCtx.String("year"),
				Genre: cCtx.String("genre"),
			})

			if err != nil {
				return err
			}

			helpers.PrintMoviesTable(movies)

			return nil
		},
	}
}

func movieCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a new movie",
		Flags: []cli.Flag{
			movieFlags["title"],
			movieFlags["released"],
			movieFlags["countryId"],
			movieFlags["personId"],
			movieFlags["genreIds"],
		},
		Action: func(cCtx *cli.Context) (err error) {
			var genres []entities.Genre

			director, country, genres, err := validateMovieRequest(
				cCtx.Int("personId"),
				cCtx.Int("countryId"),
				cCtx.IntSlice("genreIds"),
			)

			if err != nil {
				return err
			}

			movie := entities.Movie{
				Title:      cCtx.String("title"),
				ReleasedAt: cCtx.String("released"),
				Director:   director,
				Country:    country,
				Genres:     genres,
			}

			err = movie.Create()
			if err != nil {
				return err
			}

			helpers.PrintMoviesTable([]entities.Movie{movie})

			return nil
		},
	}
}

func movieUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a movie",
		Flags: []cli.Flag{
			movieFlags["id"],
			movieFlags["title"],
			movieFlags["released"],
			movieFlags["countryId"],
			movieFlags["personId"],
			movieFlags["genreIds"],
		},
		Action: func(cCtx *cli.Context) (err error) {
			var genres []entities.Genre
			movie := entities.Movie{Id: cCtx.Int("id")}

			err = movie.Find()
			if err != nil {
				fmt.Println("no movie updated")
				return nil
			}

			_ = movie.LoadDirector()
			_ = movie.LoadGenre()
			_ = movie.LoadCountry()

			personId := movie.Director.Id
			countryId := movie.Country.Id
			var genreIds []int

			for _, genre := range movie.Genres {
				genreIds = append(genreIds, genre.Id)
			}

			if cCtx.Int("personId") != 0 {
				personId = cCtx.Int("personId")
			}

			if cCtx.Int("countryId") != 0 {
				countryId = cCtx.Int("countryId")
			}

			if len(cCtx.IntSlice("genreIds")) > 0 {
				genreIds = cCtx.IntSlice("genreIds")
			}

			director, country, genres, err := validateMovieRequest(personId, countryId, genreIds)

			if err != nil {
				return err
			}

			if cCtx.String("title") != "" {
				movie.Title = cCtx.String("title")
			}

			if cCtx.String("released") != "" {
				movie.ReleasedAt = cCtx.String("released")
			}

			movie.Director = director
			movie.Country = country
			movie.Genres = genres

			err = movie.Update()
			if err != nil {
				return err
			}

			helpers.PrintMoviesTable([]entities.Movie{movie})

			return nil
		},
	}
}

func validateMovieRequest(
	directorId int,
	countryId int,
	genresIds []int,
) (director entities.Person, country entities.Country, genres []entities.Genre, err error) {
	director.Id = directorId
	err = director.Find()
	if err != nil {
		return entities.Person{}, entities.Country{}, genres, errors.New("no person found")
	}

	country.Id = countryId
	err = country.Find()
	if err != nil {
		return entities.Person{}, entities.Country{}, genres, errors.New("no country found")
	}

	for _, id := range genresIds {
		genre := entities.Genre{Id: id}
		err = genre.Find()

		if err != nil {
			return entities.Person{}, entities.Country{}, []entities.Genre{}, errors.New("no genre found")
		} else {
			genres = append(genres, genre)
		}
	}

	return director, country, genres, err
}

func movieDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a movies",
		Flags: []cli.Flag{
			movieFlags["id"],
		},
		Action: func(cCtx *cli.Context) error {
			movie := entities.Movie{Id: cCtx.Int("id")}
			err := movie.Find()

			if err != nil {
				fmt.Println("No movie deleted")
				return nil
			}

			err = movie.Delete()

			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Movie with ID %d deleted", movie.Id))

			return err
		},
	}
}
