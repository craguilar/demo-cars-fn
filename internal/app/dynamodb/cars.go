package internal

import "database/sql"

type DB struct {
	*sql.DB
}

type Tx struct {
	*sql.Tx
}
