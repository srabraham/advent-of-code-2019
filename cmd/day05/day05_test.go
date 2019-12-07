package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunPart2(t *testing.T) {
	assert.Equal(t, int64(7616021), RunPart2("input05-1.txt"))
}
