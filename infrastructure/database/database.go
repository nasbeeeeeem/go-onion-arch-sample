// Infrastructure
package database

import (
	"context"
	"go-onion-arch-sample/ent"
)

type DBClient struct {
	Client *ent.Client
}

func NewDBClient() (*DBClient, error) {
	dsn := "host=localhost port=<5432 user=postgres dbname=db password=root"
	db, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return &DBClient{Client: db}, nil
}
