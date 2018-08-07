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
	return errstack.WrapAsInfF(err, "Can't get %q from DB", title)
}

// ErrNotNoRows check if errors is not Nil and is not ErrNoRows
func ErrNotNoRows(title string, err error) errstack.E {
	if err == nil || err == pg.ErrNoRows {
		return nil
	}
	return errstack.WrapAsInfF(err, "Can't select %q from DB", title)
}
