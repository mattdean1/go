package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var input chan int
	var predicate func(int) bool

	output, err := Filter(input, predicate)

	if err != nil {
		t.Error(err)
	}

	firstOut := <-output
	if firstOut != 1 {
		t.Errorf("Filter failed, got %d, want %d", firstOut, 1)
	}
}