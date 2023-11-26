package commands

import (
	"anggafirdaus10/agung/entities"
	"anggafirdaus10/agung/helpers"
	"anggafirdaus10/agung/repositories"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	countryFlag = map[string]cli.Flag{
		"id":   &cli.IntFlag{Name: "id"},
		"name": &cli.StringFlag{Name: "name"},
	}
	CountryCommands = []*cli.Command{
		countryFindCommand(),
		countrySearchCommand(),
		countryCreateCommand(),
		countryUpdateCommand(),
		countryDeleteCommand(),
	}
)

func countryFindCommand() *cli.Command {
	return &cli.Command{
		Name:  "find",
		Usage: "find a country by id",
		Flags: []cli.Flag{
			countryFlag["id"],
		},
		Action: func(cCtx *cli.Context) error {
			entity := entities.Country{Id: cCtx.Int("id")}
			err := entity.Find()
			var entityList []entities.Country

			if err == nil {
				entityList = append(entityList, entity)
			}

			helpers.PrintCountryTable(entityList)

			return nil
		},
	}
}

func countrySearchCommand() *cli.Command {
	return &cli.Command{
		Name:  "search",
		Usage: "search countries",
		Flags: []cli.Flag{
			countryFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			countries, err := repositories.CountryRepository.
				Search(repositories.CountrySearchRequest{Name: cCtx.String("name")})

			if err != nil {
				return err
			}

			helpers.PrintCountryTable(countries)

			return nil
		},
	}
}

func countryCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a new country",
		Flags: []cli.Flag{
			countryFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			country := entities.Country{Name: cCtx.String("name")}
			err := country.Create()

			if err != nil {
				return err
			}

			helpers.PrintCountryTable([]entities.Country{country})

			return nil
		},
	}
}

func countryUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a country",
		Flags: []cli.Flag{
			countryFlag["id"],
			countryFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			country := entities.Country{Id: cCtx.Int("id")}
			err := country.Find()

			if err != nil {
				fmt.Println("No country updated")
				return nil
			}

			if cCtx.String("name") != "" {
				country.Name = cCtx.String("name")
				err = country.Update()
			}

			if err != nil {
				return err
			}

			helpers.PrintCountryTable([]entities.Country{country})

			return err
		},
	}
}

func countryDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a country",
		Flags: []cli.Flag{
			countryFlag["id"],
		},
		Action: func(cCtx *cli.Context) error {
			country := entities.Country{Id: cCtx.Int("id")}
			err := country.Find()

			if err != nil {
				fmt.Println("No country deleted")
				return nil
			}

			err = country.Delete()

			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Country with ID %d deleted", country.Id))

			return err
		},
	}
}
