package pgt

import (
	"database/sql/driver"
	"fmt"
	"strings"

	bat "github.com/robert-zaremba/go-bat"
)

// Ints is a slice of long integers for valuer interface
type Ints []int64

// Scan implements scan methods for scanner
func (ls *Ints) Scan(src interface{}) error {
	bs, err := bat.UnsafeToBytes(src)
	if err != nil {
		return err
	}
	*ls, err = ParseInt64Array(bs)
	return err
}

// Value is the valuer for integer slice
func (ls Ints) Value() (driver.Value, error) {
	res := make([]string, len(ls))
	for i, v := range ls {
		res[i] = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(res, ",")), nil
}

// Float64s is a slice of floats for valuer interface
type Float64s []float64

// Scan implements scan methods for scanner
func (f *Float64s) Scan(src interface{}) error {
	bs, err := bat.UnsafeToBytes(src)
	if err != nil {
		return err
	}
	*f, err = ParseFloatArray(bs)
	return err
}

// Value is the valuer for float slice
func (f Float64s) Value() (driver.Value, error) {
	ff := []float64(f)
	res := make([]string, len(ff))
	for i, v := range ff {
		res[i] = fmt.Sprintf("%.2f", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(res, ",")), nil
}
