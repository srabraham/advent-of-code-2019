package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(273814), Part1("input07-1.txt"))
	assert.Equal(t, int64(34579864), Part2("input07-1.txt"))
}
