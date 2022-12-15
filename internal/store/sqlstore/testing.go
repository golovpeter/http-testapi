package sqlstore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestDB(t *testing.T, databaseURL string) (*sqlx.DB, func(...string)) {
	t.Helper()

	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
