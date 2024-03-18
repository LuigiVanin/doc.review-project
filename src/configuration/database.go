package configuration

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"fmt"
)

func NewDatabase(config Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.Get("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	return db
}
