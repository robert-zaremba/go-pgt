package pgt

import (
	"database/sql/driver"
	"math/big"

	bat "github.com/robert-zaremba/go-bat"
)

// BigInt represents Postgresql numeric type for natural number
type BigInt struct {
	*big.Int
	//	Valid   bool // Valid is true if Float64 is not NULL
}

// Scan implements sql.Sanner interface
func (dst *BigInt) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	dst.Int = new(big.Int)
	s, err := bat.UnsafeToString(src)
	if err != nil {
		return err
	}
	dst.Int.SetString(s, 10)
	return nil
}

// Value implements sql/driver.Valuer
func (dst BigInt) Value() (driver.Value, error) {
	return dst.String(), nil
}
