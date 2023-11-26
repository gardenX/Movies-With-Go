package entities

import (
	"anggafirdaus10/agung/conf"
	"fmt"
	"time"
)

type Person struct {
	Id   int
	Name string
}

func (p *Person) Find() (err error) {
	err = conf.Db.
		QueryRow("SELECT id, name FROM people WHERE id = $1 AND deleted_at is NULL", p.Id).
		Scan(&p.Id, &p.Name)

	return
}

func (p *Person) Create() (err error) {
	err = conf.Db.QueryRow(fmt.Sprintf(
		"INSERT INTO people (name, created_at, updated_at) VALUES ('%[1]s', '%[2]s', '%[2]s') RETURNING id",
		p.Name,
		time.Now().Format(time.RFC3339),
	)).Scan(&p.Id)

	return err
}

func (p *Person) Update() error {
	return conf.Db.QueryRow(
		"UPDATE people SET name = $1 WHERE id = $2 AND deleted_at IS NULL RETURNING name",
		p.Name,
		p.Id,
	).Scan(&p.Name)
}

func (p *Person) Delete() error {
	return conf.Db.QueryRow(
		"UPDATE people SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL RETURNING id",
		time.Now().Format(time.RFC3339),
		p.Id,
	).Scan(&p.Id)
}
