package agung

import (
	"anggafirdaus10/agung/commands"
	"anggafirdaus10/agung/conf"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func App() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	err = conf.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	app := &cli.App{
		Name:  "Movies :",
		Usage: "Search and store Movies!",
		Action: func(*cli.Context) error {
			fmt.Println("Movies with go")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:        "movies",
				Usage:       "manage movies",
				Subcommands: commands.MovieCommands,
			},
			{
				Name:        "genres",
				Usage:       "manage genres",
				Subcommands: commands.GenreCommands,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
