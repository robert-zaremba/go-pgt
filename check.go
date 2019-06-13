package pgt

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/robert-zaremba/errstack"
)

// CheckRowsAffected asserts that expected number of rows has been affected in the
// SQL operation.
func CheckRowsAffected(title string, expected int, res orm.Result, err error) errstack.E {
	if err != nil {
		return errstack.WrapAsDomainF(err, "Query execution error %q", title)
	}
	if res.RowsAffected() != expected {
		return errstack.NewDomainF("Expected to affect %d rows, got: %d", expected,
			res.RowsAffected())
	}
	return nil
}

// CheckPgNoRows wraps pg error into errstack.E
func CheckPgNoRows(title string, err error) errstack.E {
	if err == pg.ErrNoRows {
		return errstack.WrapAsReqF(err, "Can't get %q from DB", title)
	}
	return errstack.WrapAsIOf(err, "Can't get %q from DB", title)
}

// ErrNotNoRows if errors is not Nil and is not ErrNoRows then it returns nil.
// Otherwise it will wrap the error as an IO error.
func ErrNotNoRows(title string, err error) errstack.E {
	if err == nil || err == pg.ErrNoRows {
		return nil
	}
	return errstack.WrapAsIOf(err, "Can't execute SELECT %q from DB", title)
}
