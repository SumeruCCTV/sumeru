package errors

import (
	"errors"
	"github.com/jackc/pgconn"
)

const PgErrDuplicateEntry = "23505"

func IsPgErr(err error, code string) bool {
	var e *pgconn.PgError
	return errors.As(err, &e) && e.Code == code
}
