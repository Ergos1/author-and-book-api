package psql

import "errors"

var ErrDatabaseAlreadyClosed = errors.New("database is already closed")
