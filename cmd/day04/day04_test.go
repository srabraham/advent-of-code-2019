package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_matchesPart1Criteria(t *testing.T) {
	assert.Equal(t, true, matchesPart1Criteria(11223344))
	assert.Equal(t, false, matchesPart1Criteria(11223342))
	assert.Equal(t, true, matchesPart1Criteria(11234))

}
