// Infrastructure
package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	Client *gorm.DB
}

func NewDBClient() (*DBClient, error) {
	dsn := "host=localhost user=postgres password=root dbname=db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBClient{Client: db}, nil
}
