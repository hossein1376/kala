package api

import (
	"fmt"

	_ "github.com/lib/pq"
	"kala/internal/ent"
	"kala/internal/structure"
)

func openSqlDB(cfg structure.Config) (*ent.Client, error) {
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
