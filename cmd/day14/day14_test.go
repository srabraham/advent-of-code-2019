package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 158482, Part1("input14-1.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 7993831, Part2("input14-1.txt"))
}
