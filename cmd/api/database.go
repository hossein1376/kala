package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/hossein1376/kala/config"
)

func openSqlDB(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Sql.Host,
		cfg.DB.Sql.Port,
		cfg.DB.Sql.Username,
		cfg.DB.Sql.Password,
		cfg.DB.Sql.Name)

	client, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

func openRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.DB.Redis.Host, cfg.DB.Redis.Port),
		Password: cfg.DB.Redis.Password,
		DB:       cfg.DB.Redis.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed establishing Redis connection, %v", err)
	}

	_, err = rdb.FlushAll(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to flush Redis db, %v", err)
	}

	return rdb, nil
}
