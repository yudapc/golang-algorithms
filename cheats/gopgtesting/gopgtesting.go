package gopgtesting

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func NewDBConnection() *pg.DB {
	godotenv.Load(".env")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
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
