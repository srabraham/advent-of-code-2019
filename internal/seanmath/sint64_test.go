package seanmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, int64(100), Max([]int64{100, 1, 2, -10}))
	assert.Equal(t, int64(100), Max([]int64{1, 2, 100, -10}))
}

func TestMin(t *testing.T) {
	assert.Equal(t, int64(100), Min([]int64{100, 101, 200, 1000}))
	assert.Equal(t, int64(100), Min([]int64{1000, 200, 100, 10000}))
}

func TestParseString(t *testing.T) {
	assert.Equal(t, int64(-9001), ParseString(ToString(-9001)))
}