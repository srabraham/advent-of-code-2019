package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1747, Part1("input11-1.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, ""+
		`. # # # # . . # # . . . # # . . # # # . . # . . # . # . . # . # . . . . # # # . . . . 
. . . . # . # . . # . # . . # . # . . # . # . . # . # . # . . # . . . . # . . # . . . 
. . . # . . # . . . . # . . . . # . . # . # # # # . # # . . . # . . . . # # # . . . . 
. . # . . . # . . . . # . # # . # # # . . # . . # . # . # . . # . . . . # . . # . . . 
. # . . . . # . . # . # . . # . # . # . . # . . # . # . # . . # . . . . # . . # . . . 
. # # # # . . # # . . . # # # . # . . # . # . . # . # . . # . # # # # . # # # . . . . 
`, Part2("input11-1.txt"))
}
