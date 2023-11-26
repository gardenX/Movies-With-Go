package commands

import (
	"anggafirdaus10/agung/entities"
	"anggafirdaus10/agung/helpers"
	"anggafirdaus10/agung/repositories"
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	personFlag = map[string]cli.Flag{
		"id":   &cli.IntFlag{Name: "id"},
		"name": &cli.StringFlag{Name: "name"},
	}
	PersonCommands = []*cli.Command{
		personFindCommand(),
		personSearchCommand(),
		personCreateCommand(),
		personUpdateCommand(),
		personDeleteCommand(),
	}
)

func personFindCommand() *cli.Command {
	return &cli.Command{
		Name:  "find",
		Usage: "find a person by id",
		Flags: []cli.Flag{
			personFlag["id"],
		},
		Action: func(cCtx *cli.Context) error {
			entity := entities.Person{Id: cCtx.Int("id")}
			err := entity.Find()
			var entityList []entities.Person

			if err == nil {
				entityList = append(entityList, entity)
			}

			helpers.PrintPeopleTable(entityList)

			return nil
		},
	}
}

func personSearchCommand() *cli.Command {
	return &cli.Command{
		Name:  "search",
		Usage: "search people",
		Flags: []cli.Flag{
			personFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			people, err := repositories.PersonRepository.
				Search(repositories.PersonSearchRequest{Name: cCtx.String("name")})

			if err != nil {
				return err
			}

			helpers.PrintPeopleTable(people)

			return nil
		},
	}
}

func personCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a new person",
		Flags: []cli.Flag{
			personFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			person := entities.Person{Name: cCtx.String("name")}
			err := person.Create()

			if err != nil {
				return err
			}

			helpers.PrintPeopleTable([]entities.Person{person})

			return nil
		},
	}
}

func personUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a person",
		Flags: []cli.Flag{
			personFlag["id"],
			personFlag["name"],
		},
		Action: func(cCtx *cli.Context) error {
			person := entities.Person{Id: cCtx.Int("id")}
			err := person.Find()

			if err != nil {
				fmt.Println("No person updated")
				return nil
			}

			if cCtx.String("name") != "" {
				person.Name = cCtx.String("name")
				err = person.Update()
			}

			if err != nil {
				return err
			}

			helpers.PrintPeopleTable([]entities.Person{person})

			return err
		},
	}
}

func personDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a person",
		Flags: []cli.Flag{
			personFlag["id"],
		},
		Action: func(cCtx *cli.Context) error {
			person := entities.Person{Id: cCtx.Int("id")}
			err := person.Find()

			if err != nil {
				fmt.Println("No person deleted")
				return nil
			}

			err = person.Delete()

			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Person with ID %d deleted", person.Id))

			return err
		},
	}
}
