package pgt

import (
	"math/rand"

	"github.com/elgs/gostrgen"
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type StringSuite struct{}

func (suite *StringSuite) TestStringsContainsAllSorted(c *C) {
	var s, subset Strings

	// 1. Check empty slices
	c.Assert(s.ContainsAllSorted(nil), IsTrue)
	c.Assert(s.ContainsAllSorted(subset), IsTrue)

	// 2. Every slice should contain an empty slice
	s = Strings{"a1", "a2", "b1"}
	c.Assert(s.ContainsAllSorted(nil), IsTrue)
	c.Assert(s.ContainsAllSorted(subset), IsTrue)

	// 3. Empty slice can't be a superset of non-empty one
	c.Assert(subset.ContainsAllSorted(s), IsFalse, Comment("Empty slice can't be a super set of non-empty"))

	// 4. Check non empty slices
	c.Assert(s.ContainsAllSorted(Strings{"b1"}), IsTrue)
	c.Assert(s.ContainsAllSorted(Strings{"a1", "b1"}), IsTrue)
	c.Assert(s.ContainsAllSorted(s), IsTrue)

	c.Assert(s.ContainsAllSorted(Strings{"a1", "a3", "b1"}), IsFalse)
	c.Assert(s.ContainsAllSorted(Strings{"aa", "b12"}), IsFalse)
	c.Assert(s.ContainsAllSorted(Strings{"a1", "a2", "b1", "b2"}), IsFalse)
}

func (suite *StringSuite) TestStringForEachPair(c *C) {
	testCases := []struct {
		Input  Strings
		Result [][2]string
	}{
		{
			Input:  nil,
			Result: [][2]string{},
		},
		{
			Input:  Strings{"1"},
			Result: [][2]string{},
		},
		{
			Input:  Strings{"1", "2"},
			Result: [][2]string{{"1", "2"}},
		},
		{
			Input:  Strings{"1", "2", "3"},
			Result: [][2]string{{"1", "2"}, {"2", "3"}},
		},
		{
			Input:  Strings{"1", "2", "3", "4"},
			Result: [][2]string{{"1", "2"}, {"2", "3"}, {"3", "4"}},
		},
	}
	for _, testCase := range testCases {
		received := [][2]string{}
		listener := func(a, b string) {
			received = append(received, [2]string{a, b})
		}
		testCase.Input.ForEachPair(listener)
		c.Assert(testCase.Result, DeepEquals, received)
	}
}

func (suite *StringSuite) TestStringJSON(c *C) {
	testCases := []String{
		{Valid: false},
		{Valid: true},
		{String: "\"\"", Valid: true},
		{String: "null", Valid: true},
		{String: "\"null\"", Valid: true},
	}

	for i, testCase := range testCases {
		comment := Commentf("Error when testing case %d of pgu.String on %v", i, testCase)
		b, err := testCase.MarshalJSON()
		c.Assert(err, IsNil, comment)
		var received String
		err = received.UnmarshalJSON(b)
		c.Assert(err, IsNil, comment)
		c.Assert(testCase, DeepEquals, received)
	}
}

func (suite *StringSuite) TestStringsDistinct(c *C) {
	testCases := []struct{ input, output Strings }{
		{input: nil, output: nil},
		{input: Strings{}, output: Strings{}},
		{input: Strings{"a"}, output: Strings{"a"}},
		{input: Strings{"a", "a"}, output: Strings{"a"}},
		{input: Strings{"a", "b", "a"}, output: Strings{"a", "b"}},
		{input: Strings{"a", "b", "a", "c"}, output: Strings{"a", "b", "c"}},
		{input: Strings{"a", "b", "", "a", "", "c"}, output: Strings{"a", "b", "", "c"}},
	}

	for _, testCase := range testCases {
		c.Assert(testCase.input.Distinct(), DeepEquals, testCase.output)
	}
}

func (suite *StringSuite) checkStrArray(c *C, ss Strings) {
	comment := Commentf("Error when testing pgu.Strings on %v", ss)
	s, err := ss.Value()
	c.Assert(err, IsNil, comment)
	var result Strings
	err = result.Scan([]byte(s.(string)))
	c.Assert(err, IsNil, comment)
	c.Assert(result, DeepEquals, ss, comment)
}

func (suite *StringSuite) TestArraySameAfterMarshallingAndUmarshaling(c *C) {
	extraChars := "{}<>()'`,.;\"///\\\\\\\n\t "
	for t := 0; t < 3000; t++ {
		for elems := 1; elems < 20; elems++ {
			suite.checkStrArray(c,
				Strings(generateStrings(c, elems, extraChars)))
		}
	}
}

func (suite *StringSuite) TestVal(c *C) {
	testVal(c, `{"mary"}`, "mary")
	testVal(c, "{\"\n\"}", "\n")
	testVal(c, "{\"\n\"}", `
`)
	testVal(c, `{"mary","had"}`, "mary", "had")
	testVal(c, `{}`)
	testVal(c, `{"\""}`, `"`)
	testVal(c, `{"\\"}`, `\`)
	testVal(c, "{\"a\nb\"}", "a\nb")
}

func testVal(c *C, expected string, ss ...string) {
	comment := Commentf("Error when testing pgu.Strings on %v", ss)
	s, err := Strings(ss).Value()
	c.Assert(err, IsNil, comment)
	c.Assert(s, Equals, expected)
}

func (suite *StringSuite) TestWrongInputDoesNotCausePanic(c *C) {
	badInput := []string{`{"}`, `{""a}`, `{mary,"}`, `{mary,}`}
	for _, b := range badInput {
		v, err := parseArray(b)
		c.Assert(err, NotNil, Commentf("Expected error for %s, but parsed as %v", b, v))
	}

}

func generateStrings(c *C, size int, extraCharset string) []string {
	var ls []string
	for i := 0; i < size; i++ {
		l := rand.Intn(16) + 1
		s, err := gostrgen.RandGen(l, gostrgen.All, extraCharset, "")
		c.Assert(err, IsNil)
		ls = append(ls, s)
	}
	return ls
}
