package entities

import (
	"anggafirdaus10/agung/conf"
	"fmt"
	"time"
)

type Country struct {
	Id   int
	Name string
}

func (c *Country) Find() (err error) {
	return conf.Db.
		QueryRow("SELECT id, name FROM countries WHERE id = $1 AND deleted_at is NULL", c.Id).
		Scan(&c.Id, &c.Name)
}

func (c *Country) Create() (err error) {
	return conf.Db.QueryRow(fmt.Sprintf(
		"INSERT INTO countries (name, created_at, updated_at) VALUES ('%[1]s', '%[2]s', '%[2]s') RETURNING id",
		c.Name,
		time.Now().Format(time.RFC3339),
	)).Scan(&c.Id)
}

func (c *Country) Update() error {
	return conf.Db.QueryRow(
		"UPDATE countries SET name = $1 WHERE id = $2 AND deleted_at IS NULL RETURNING name",
		c.Name,
		c.Id,
	).Scan(&c.Name)
}

func (c *Country) Delete() error {
	return conf.Db.QueryRow(
		"UPDATE countries SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL RETURNING id",
		time.Now().Format(time.RFC3339),
		c.Id,
	).Scan(&c.Id)
}
