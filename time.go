package pgt

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"strings"
	"time"

	bat "github.com/robert-zaremba/go-bat"
)

// Time overrides time.Time for two reasons:
// (1) json serialization
// (2) nullability.
//
// On default time.Time serializes two json as string (ISO format).
// For communication with client
// we need serialize to UNIX timestamp.
//
// As time.Time is as struct, we are unable to distinguish null value
// (best we can do is use time.Time.IsZero()).
// To overcome this issue, agtime.Time has `Valid` field which enables
// us to handle nulls property. Mainly replicates pq.NullTime behavior
type Time struct {
	time.Time
	Valid bool
}

// UTCNow is an utility method for creating Time
// representing "now" in UTC time zone
func UTCNow() Time {
	return Time{time.Now().UTC(), true}
}

// NewTime creates new valid UTC Time
func NewTime(t time.Time) Time {
	if t.Location() != time.UTC {
		return Time{t.UTC(), true}
	}
	return Time{t, true}
}

// Scan implements Scanner interface
func (t *Time) Scan(value interface{}) error {
	t.Time, t.Valid = value.(time.Time)
	t.Time = t.Time.UTC()
	return nil
}

// Value implements Valuer interface
func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time.UTC(), nil
}

// MarshalJSON implements Marshaller interface
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nullbytes, nil
	}
	return []byte(bat.I64toa(t.Unix())), nil
}

// UnmarshalJSON implements Unmarshaller interface
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, nullbytes) {
		return nil
	}
	asString := strings.Trim(string(data), "\"")
	// used when aggregated columns are returned from db
	parsed, err := time.Parse("2006-01-02T15:04:05", asString)
	if err == nil {
		*t = Time{parsed, true}
		return nil
	}
	unixTime, err := bat.Atoi64(asString)
	if err != nil {
		return err
	}
	// This is new for 0 or zero time
	if unixTime <= 0 {
		return nil
	}
	*t = Time{time.Unix(unixTime, 0).UTC(), true}
	return nil
}

// MarshalBinary implements binary encoding for time
// This pair of methods are used if agtime.Time is msgpacked.
//
func (t Time) MarshalBinary() ([]byte, error) {
	bs := make([]byte, 16)
	binary.PutVarint(bs, t.Unix())
	binary.PutVarint(bs[8:], int64(t.Nanosecond()))
	return bs, nil
}

// UnmarshalBinary implements binary decoding for time
func (t *Time) UnmarshalBinary(data []byte) (err error) {
	var unix, nano int64
	if unix, err = getBytes(data); err != nil {
		return
	}
	if nano, err = getBytes(data[8:]); err != nil {
		return
	}
	*t = Time{time.Unix(unix, nano).UTC(), true}
	return nil
}

// Add is a proxy for the time:Time.Add method
func (t Time) Add(d time.Duration) Time {
	return Time{t.Time.Add(d), t.Valid}
}

func getBytes(data []byte) (int64, error) {
	x, n := binary.Varint(data)
	if n == 0 {
		return x, errors.New("Can't unmarshal integer: buffer too small")
	} else if n < 0 {
		return x, errors.New("Can't unmarshal integer: value larger than 64 bits (overflow)")
	}
	return x, nil
}
