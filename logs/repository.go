package logs

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // repository assumes sqlite
)

type Repository struct {
	db *sqlx.DB
}

type Game struct {
	Day              string
	ID               int
	Timestamp        time.Time
	Size             int
	Player1, Player2 string
	Result           string
	Winner           string
	Moves            int
}

func Open(db string) (*Repository, error) {
	sql, err := sqlx.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&mode=rwc", db))
	if err != nil {
		return nil, err
	}
	return &Repository{db: sql}, nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) DB() *sqlx.DB {
	return r.db
}
