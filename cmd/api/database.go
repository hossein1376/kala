package api

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/hossein1376/kala/cmd"
	"github.com/hossein1376/kala/internal/ent"
)

func openSqlDB(cfg *cmd.Config) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Sql.Host,
		cfg.DB.Sql.Port,
		cfg.DB.Sql.Username,
		cfg.DB.Sql.Password,
		cfg.DB.Sql.Name)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
