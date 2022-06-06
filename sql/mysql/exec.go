package mysql

import (
	"database/sql"
	"errors"
	"github.com/mono83/query"
)

// Query method executes given query on given database.
func Query(db *sql.DB, q query.Query) (*sql.Rows, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	stmt, err := QueryToStatement(q)
	if err != nil {
		return nil, err
	}

	return db.Query(stmt.Query(), stmt.Args()...)
}
