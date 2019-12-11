package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 221, Part1("input10-1.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, GridPos{x: 8, y: 6}, Part2("input10-1.txt", 11, 11, 200))
}
