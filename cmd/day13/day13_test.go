package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	// Part 1
	assert.Equal(t, 200, RunGame("input13-0.txt", false).count(Block))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(9803), RunGame("input13-1.txt", false).score)
}
