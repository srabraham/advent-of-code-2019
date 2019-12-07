package intcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunIntCodeWithChannels(t *testing.T) {
	cmds := []int64{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	aIn := make(chan int64, 1)
	aOut := make(chan int64, 1)
	aIn <- 9
	result := RunIntCodeWithChannels(cmds, aIn, aOut)
	assert.Equal(t, int64(1001), result)
	assert.Equal(t, int64(1001), <-aOut)
}
