package pgt

import (
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
