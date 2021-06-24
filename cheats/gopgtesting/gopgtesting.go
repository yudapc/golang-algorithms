package gopgtesting

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func NewDBConnection() *pg.DB {
	username := "postgres"
	password := "12345678"
	port := "5432"
	database := "gopgtesting"
	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", username, password, port, database)
	opt, err := pg.ParseURL(dbUrl)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return db
}
