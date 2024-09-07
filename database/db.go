package database

import (
	"database/sql"

	"github.com/fermyon/spin/sdk/go/v2/sqlite"
)

func New() *sql.DB {
	return sqlite.Open("default")
}
