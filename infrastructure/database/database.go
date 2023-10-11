// Infrastructure
package database

import (
	"context"
	"go-onion-arch-sample/ent"

	_ "github.com/lib/pq" // This line is required to register the Postgres driver.

	"entgo.io/ent/dialect"
)

type DBClient struct {
	Client *ent.Client
}

func NewDBClient() (*DBClient, error) {
	dsn := "host=localhost port=5432 user=postgres dbname=db password=root sslmode=disable"
	db, err := ent.Open(
		dialect.Postgres,
		dsn,
	)
	if err != nil {
		return nil, err
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	return &DBClient{Client: db}, nil
}
