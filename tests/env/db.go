package env

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type TestDb struct {
	*sql.DB
}

func InitTestDb(dsn string) (*TestDb, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &TestDb{
		db,
	}, nil
}

func (db *TestDb) Up() error {
	if err := goose.Up(db.DB, "../../migrations"); err != nil {
		return err
	}
	return nil
}
func (db *TestDb) Down() error {
	if err := goose.Down(db.DB, "../../migrations"); err != nil {
		return err
	}
	return nil
}
