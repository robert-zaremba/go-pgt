package pgt

import (
	"time"

	. "github.com/robert-zaremba/checkers"
	. "gopkg.in/check.v1"
)

type TimeSuite struct{}

func (suite *TimeSuite) TestTimeMarshalBinary(c *C) {
	t := UTCNow()
	b, err := t.MarshalBinary()
	c.Assert(err, IsNil)

	var t2 = new(Time)
	err = t2.UnmarshalBinary(b)
	c.Assert(err, IsNil)
	c.Assert(t.Time, Equals, t2.Time)
}

func (suite *TimeSuite) TestTimeMarshalJSON(c *C) {
	t := UTCNow()
	d, err := t.MarshalJSON()
	c.Assert(err, IsNil)

	var t2 = new(Time)
	err = t2.UnmarshalJSON(d)
	c.Assert(err, IsNil)
	// Marshall marshals up to second accuracy
	c.Assert(t.Time, WithinDuration, t2.Time, time.Second)
	c.Assert(t2.Valid, IsTrue)

	err = t2.UnmarshalJSON(nullbytes)
	c.Assert(err, IsNil)
	c.Assert(t2.Valid, IsFalse)
}
