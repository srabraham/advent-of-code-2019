package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, []int64{3906448201}, RunBoost("input09-1.txt", []int64{1}))
	assert.Equal(t, []int64{59785}, RunBoost("input09-1.txt", []int64{2}))
}
