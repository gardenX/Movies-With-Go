package entities

import (
	"anggafirdaus10/agung/conf"
	"fmt"
	"time"
)

type Genre struct {
	Id   int
	Name string
}

func (g *Genre) Find() (err error) {
	return conf.Db.
		QueryRow("SELECT id, name FROM genres WHERE id = $1 AND deleted_at is NULL", g.Id).
		Scan(&g.Id, &g.Name)
}

func (g *Genre) Create() (err error) {
	return conf.Db.QueryRow(fmt.Sprintf(
		"INSERT INTO genres (name, created_at, updated_at) VALUES ('%[1]s', '%[2]s', '%[2]s') RETURNING id",
		g.Name,
		time.Now().Format(time.RFC3339),
	)).Scan(&g.Id)
}

func (g *Genre) Update() error {
	return conf.Db.QueryRow(
		"UPDATE genres SET name = $1 WHERE genres.id = $2 AND deleted_at IS NULL RETURNING genres.name",
		g.Name,
		g.Id,
	).Scan(&g.Name)
}

func (g *Genre) Delete() error {
	return conf.Db.QueryRow(
		"UPDATE genres SET deleted_at = $1 WHERE genres.id = $2 AND deleted_at IS NULL RETURNING genres.id",
		time.Now().Format(time.RFC3339),
		g.Id,
	).Scan(&g.Id)
}
