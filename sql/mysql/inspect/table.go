package inspect

import (
	"database/sql"
	"errors"
	"github.com/mono83/query"
	"github.com/mono83/query/fields"
)

type virtualColumnType struct {
	dataType string
	nullable bool
}

func (v virtualColumnType) DatabaseTypeName() string { return v.dataType }
func (v virtualColumnType) Nullable() (bool, bool)   { return v.nullable, true }

// TableFields analyzes table structure and returns slice of
// sortable/filterable fields.
func TableFields(db *sql.DB, database, table string) ([]query.Field, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	rows, err := db.Query(
		"SELECT `COLUMN_NAME`, `DATA_TYPE`, `COLUMN_KEY`, `IS_NULLABLE` FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA`=? AND `TABLE_NAME`=?",
		database, table,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []query.Field
	var columnName, columnDataType, columnKey, isNull string
	for rows.Next() {
		if err := rows.Scan(&columnName, &columnDataType, &columnKey, &isNull); err != nil {
			return nil, err
		}

		if notIgnored(columnDataType) {
			// Data type
			dt := virtualColumnType{
				dataType: columnDataType,
				nullable: isNull == "YES",
			}

			// Constructing field
			if len(columnKey) > 0 {
				out = append(out, fields.New(columnName, true, isSortable(columnDataType), dt))
			} else {
				out = append(out, fields.New(columnName, false, false, dt))
			}
		}
	}
	return out, nil
}

// TablesFields analyzes all tables within database and returns
// slice of  sortable/filterable fields.
func TablesFields(db *sql.DB, database string) (map[string][]query.Field, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	rows, err := db.Query(
		"SELECT `TABLE_NAME` FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA`=?",
		database,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[string][]query.Field{}
	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}

		fields, err := TableFields(db, database, tableName)
		if err != nil {
			return nil, err
		}
		out[tableName] = fields
	}
	return out, nil
}

func notIgnored(dataType string) bool {
	switch dataType {
	case "text", "blob":
		return false
	}
	return true
}

func isSortable(dataType string) bool {
	switch dataType {
	case "enum":
		return false // Enums are not sortable
	}
	return true
}
