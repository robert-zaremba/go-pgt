package pgt

import (
	"bytes"
	"database/sql/driver"
	"math/big"

	bat "github.com/robert-zaremba/go-bat"
)

// BigInt represents Postgresql numeric type for natural number
type BigInt struct {
	*big.Int
	//	Valid   bool // Valid is true if Float64 is not NULL
}

// NewInt allocates and returns a new Int set to x.
func NewBigInt(x int64) BigInt {
	return BigInt{big.NewInt(x)}
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
	if dst.Int == nil {
		return nil, nil
	}
	return dst.String(), nil
}

// IsNull returns tru iff the inner value is nil
func (dst BigInt) IsNull() bool {
	return dst.Int == nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (dst *BigInt) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if bytes.Equal(data, nullbytes) {
		return nil
	}
	if dst.Int == nil {
		dst.Int = new(big.Int)
	}
	return dst.Int.UnmarshalText(data)
}
