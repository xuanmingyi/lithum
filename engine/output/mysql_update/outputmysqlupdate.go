package mysqlupdate

import (
	"context"
	"database/sql"
)

const ModuleName = "mysqlupdate"

type OutputMySQLUpdate struct {
	Ctx context.Context
	DB  *sql.DB
	DSN string
}
