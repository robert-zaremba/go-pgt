package pgt

import (
	. "gopkg.in/check.v1"
)

type BigIntS struct{}

func (suite *BigIntS) TestBigIntMarshal(c *C) {
	// check null value
	var i, dest BigInt
	testMarshalJSON(i, &dest, c)
	c.Assert(i, DeepEquals, dest)

	// check real value
	i = NewBigInt(239)
	testMarshalJSON(i, &dest, c)
	c.Assert(i, DeepEquals, dest)

	// check composed value
	type composed struct {
		Number BigInt
		B      string
	}
	var obj = composed{i, "abc"}
	var destObj composed
	testMarshalJSON(obj, &destObj, c)
	c.Assert(obj, DeepEquals, destObj)
}
