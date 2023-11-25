package commands

import (
	"anggafirdaus10/agung/entities"
	"anggafirdaus10/agung/helpers"
	"anggafirdaus10/agung/repositories"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	movieFlags = map[string]cli.Flag{
		"id":    &cli.IntFlag{Name: "id"},
		"title": &cli.StringFlag{Name: "title", Aliases: []string{"t"}},
		"year":  &cli.StringFlag{Name: "year", Aliases: []string{"y"}},
		"genre": &cli.StringFlag{Name: "genre", Aliases: []string{"g"}},
	}
	MovieCommands = []*cli.Command{
		movieFindCommand(),
		movieSearchCommand(),
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
			movie, err := repositories.MovieRepository.Find(cCtx.String("id"))
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

func movieDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a movies",
		Flags: []cli.Flag{
			movieFlags["id"],
		},
		Action: func(cCtx *cli.Context) error {
			movie, err := repositories.MovieRepository.Find(cCtx.String("id"))

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
