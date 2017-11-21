package pgt

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

func testMarshalJSON(src, dest interface{}, c *C) {
	b, err := json.Marshal(src)
	c.Assert(err, IsNil)
	err = json.Unmarshal(b, dest)
	c.Assert(err, IsNil)
}
