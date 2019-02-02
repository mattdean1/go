package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var input <-chan int
	var predicate func(int) bool

	_, err := Filter(input, predicate)

	if err != nil {
		t.Errorf("fail")
	}
}