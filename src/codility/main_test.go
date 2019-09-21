package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPermMissingElem(c *C) {
	var solutions = [...]int{4, 6, 3}

	c.Assert(PermMissingElem([]int{2, 3, 1, 5}), Equals, solutions[0])
	c.Assert(PermMissingElem([]int{1, 2, 3, 4, 5, 7}), Equals, solutions[1])
}

func (s *MySuite) BenchmarkPermMissingElem(c *C) {
	for i := 0; i < c.N; i++ {
		PermMissingElem([]int{1, 2, 3, 4, 5, 7})
	}
}

func (s *MySuite) TestTapeEquilibrium(c *C) {
	c.Assert(TapeEquilibrium([]int{3, 1, 2, 4, 3}), Equals, 1)
}

func (s *MySuite) BenchmarkTapeEquilibrium(c *C) {
	for i := 0; i < c.N; i++ {
		TapeEquilibrium([]int{3, 1, 2, 4, 3})
	}
}