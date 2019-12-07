package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	if fuel(100756) != 33583 {
		t.Error("wrong mass!")
	}
}

func TestThing(t *testing.T) {
	res1, res2 := RunDay1("input01-1.txt")
	assert.Equal(t, res1, int64(3256114))
	assert.Equal(t, res2, int64(4881302))
}
