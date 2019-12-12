package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const(
	realInput = `<x=5, y=4, z=4>
<x=-11, y=-11, z=-3>
<x=0, y=7, z=0>
<x=-13, y=2, z=10>`
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 10845, Part1(realInput, 1000))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 551272644867044, Part2(realInput))
}
