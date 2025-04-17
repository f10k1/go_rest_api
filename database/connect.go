package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer conn.Close()

	return conn
}

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var Database = map[string]ClientProfile{
	"1": {
		Email: "test@test.com",
		Id:    "1",
		Name:  "Jan Kowalski",
		Token: "234",
	},
}
