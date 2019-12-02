package main

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	res := run([]int64{1,0,0,0,99})
	log.Print(res)
}