package main

import (
	"context"
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
)

func migrate(db *sqlx.DB, path string) error {
	// open directory containing migration files
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// start a transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	// loop over migration files; read and execute each one.
	// if something goes wrong, rollback the transaction and return.
	for _, file := range dir {
		content, err := os.ReadFile(path + "/" + file.Name())
		if err != nil {
			return errors.Join(err, tx.Rollback())
		}

		_, err = tx.Exec(string(content))
		if err != nil {
			return errors.Join(err, tx.Rollback())
		}
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	return nil
}
