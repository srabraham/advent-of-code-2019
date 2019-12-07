package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	noun, verb := Part2("input02-1a.txt")
	assert.Equal(t, int64(54), noun)
	assert.Equal(t, int64(85), verb)
}
