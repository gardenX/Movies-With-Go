package repositories

import (
	"anggafirdaus10/agung/conf"
	"anggafirdaus10/agung/entities"
	"fmt"
)

type PersonSearchRequest struct {
	Name string
}

func (g personRepository) Search(request PersonSearchRequest) (people []entities.Person, err error) {
	people = []entities.Person{}
	rows, err := conf.Db.Query(
		"SELECT id, name FROM people WHERE deleted_at IS NULL and name ILIKE $1",
		fmt.Sprintf("%[1]s%[2]s%[1]s", "%", request.Name),
	)

	if err != nil {
		return []entities.Person{}, err
	}

	for rows.Next() {
		person := entities.Person{}
		err = rows.Scan(&person.Id, &person.Name)

		if err != nil {
			return []entities.Person{}, err
		}

		people = append(people, person)
	}

	return people, err
}
