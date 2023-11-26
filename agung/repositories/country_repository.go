package repositories

import (
	"anggafirdaus10/agung/conf"
	"anggafirdaus10/agung/entities"
	"fmt"
)

type CountrySearchRequest struct {
	Name string
}

func (g countryRepository) Search(request CountrySearchRequest) (people []entities.Country, err error) {
	rows, err := conf.Db.Query(
		"SELECT id, name FROM countries WHERE deleted_at IS NULL and name ILIKE $1",
		fmt.Sprintf("%[1]s%[2]s%[1]s", "%", request.Name),
	)

	if err != nil {
		return people, err
	}

	for rows.Next() {
		country := entities.Country{}
		err = rows.Scan(&country.Id, &country.Name)

		if err != nil {
			return people, err
		}

		people = append(people, country)
	}

	return people, err
}
