package mssql

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func NewConnection(connString string) (*sql.DB, error) {
	return sql.Open("sqlserver", connString)
}
