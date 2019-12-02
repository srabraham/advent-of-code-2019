package main

import "testing"

func Test1(t *testing.T) {
	if fuel(100756) != 33583 {
		t.Error("wrong mass!")
	}
}