package api

import (
	"fmt"

	"kala/cmd"
	"kala/internal/ent"

	_ "github.com/lib/pq"
)

func openSqlDB(cfg cmd.Config) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Sql.Host,
		cfg.Sql.Port,
		cfg.Sql.Username,
		cfg.Sql.Password,
		cfg.Sql.Name)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
