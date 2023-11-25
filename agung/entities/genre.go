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

func (g *Genre) Create() (err error) {
	err = conf.Db.QueryRow(fmt.Sprintf(
		"INSERT INTO genres (name, created_at, updated_at) VALUES ('%[1]s', '%[2]s', '%[2]s') RETURNING id",
		g.Name,
		time.Now().Format(time.RFC3339),
	)).Scan(&g.Id)

	return err
}
