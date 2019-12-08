package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1330, Part1("input08-1.txt", 25, 6))
	// Says FAHEF
	expectedPart2 := `X X X X     X X     X     X   X X X X   X X X X   
X         X     X   X     X   X         X         
X X X     X     X   X X X X   X X X     X X X     
X         X X X X   X     X   X         X         
X         X     X   X     X   X         X         
X         X     X   X     X   X X X X   X         
`
	assert.Equal(t, expectedPart2, Part2("input08-1.txt", 25, 6))
}
