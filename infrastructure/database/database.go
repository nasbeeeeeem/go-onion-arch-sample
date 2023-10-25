// Infrastructure
package database

import (
	"context"
	"go-onion-arch-sample/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

type DBClient struct {
	Client *ent.Client
}

func NewDBClient() (*DBClient, error) {
	// vercelDsn := "postgres://default:4DrVmFQl8IZk@ep-throbbing-heart-07304335.ap-southeast-1.postgres.vercel-storage.com:5432/verceldb"
	// neonDsn := "postgres://nasbeeeeeem:YK8pIa2CbiTs@ep-dry-art-16713738.ap-southeast-1.aws.neon.tech/neondb"
	dsn := "host=localhost port=5432 user=postgres dbname=db password=root sslmode=disable"
	db, err := ent.Open(
		dialect.Postgres,
		dsn,
	)
	if err != nil {
		return nil, err
	}

	// マイグレーションの実行
	if err := db.Schema.Create(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return &DBClient{Client: db}, nil
}
