package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	inputValues := []int{1, 1, 2, 3, 5}
	input := make(chan int, 5)
	var predicate func(int) bool

	for v := range inputValues {
		input<-v
	}
	close(input)

	output, err := Filter(input, predicate)
	if err != nil {
		t.Error(err)
	}

	for v := range inputValues {
		outputValue := <-output
		if outputValue != v {
			t.Errorf("Filter failed, got %d, want %d", outputValue, v)
		}
	}
}