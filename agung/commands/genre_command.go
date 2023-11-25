package commands

import (
	"anggafirdaus10/agung/entities"
	"anggafirdaus10/agung/helpers"
	"anggafirdaus10/agung/repositories"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	genreFlags = map[string]cli.Flag{
		"id":   &cli.IntFlag{Name: "id"},
		"name": &cli.StringFlag{Name: "name"},
	}
	GenreCommands = []*cli.Command{
		genreFindCommand(),
		genreSearchCommand(),
		genreCreateCommand(),
		genreUpdateCommand(),
		genreDeleteCommand(),
	}
)

func genreFindCommand() *cli.Command {
	return &cli.Command{
		Name:  "find",
		Usage: "find a genre by id",
		Flags: []cli.Flag{
			genreFlags["id"],
		},
		Action: func(cCtx *cli.Context) error {
			genre := entities.Genre{Id: cCtx.Int("id")}
			err := genre.Find()
			var genres []entities.Genre

			if err == nil {
				genres = append(genres, genre)
			}

			helpers.PrintGenresTable(genres)

			return nil
		},
	}
}

func genreSearchCommand() *cli.Command {
	return &cli.Command{
		Name:  "search",
		Usage: "search genres",
		Flags: []cli.Flag{
			genreFlags["name"],
		},
		Action: func(cCtx *cli.Context) error {
			genres, err := repositories.GenreRepository.
				Search(repositories.GenreSearchRequest{Name: cCtx.String("name")})

			if err != nil {
				return err
			}

			helpers.PrintGenresTable(genres)

			return nil
		},
	}
}

func genreCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a new genre",
		Flags: []cli.Flag{
			genreFlags["name"],
		},
		Action: func(cCtx *cli.Context) error {
			genre := entities.Genre{Name: cCtx.String("name")}
			err := genre.Create()

			if err != nil {
				return err
			}

			helpers.PrintGenresTable([]entities.Genre{genre})

			return nil
		},
	}
}

func genreUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a genre",
		Flags: []cli.Flag{
			genreFlags["id"],
			genreFlags["name"],
		},
		Action: func(cCtx *cli.Context) error {
			genre := entities.Genre{Id: cCtx.Int("id")}
			err := genre.Find()

			if err != nil {
				fmt.Println("No genre updated")
				return nil
			}

			if cCtx.String("name") != "" {
				genre.Name = cCtx.String("name")
				err = genre.Update()
			}

			if err != nil {
				return err
			}

			helpers.PrintGenresTable([]entities.Genre{genre})

			return err
		},
	}
}

func genreDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a genre",
		Flags: []cli.Flag{
			genreFlags["id"],
		},
		Action: func(cCtx *cli.Context) error {
			genre := entities.Genre{Id: cCtx.Int("id")}
			err := genre.Find()

			if err != nil {
				fmt.Println("No genre deleted")
				return nil
			}

			err = genre.Delete()

			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Genre with ID %d deleted", genre.Id))

			return err
		},
	}
}
