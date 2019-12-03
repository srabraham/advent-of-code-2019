package intcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	res := RunIntCode([]int64{1, 1, 1, 4, 99, 5, 6, 0, 99})
	assert.Equal(t, []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}, res)
}
