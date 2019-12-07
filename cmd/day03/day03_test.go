package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunPart2(t *testing.T) {
	assert.Equal(t, 21666, RunPart2("input03-1.txt"))
}
