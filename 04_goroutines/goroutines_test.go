package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	input := make(chan int, 2)
	var predicate func(int) bool

	input<-1
	input<-2
	close(input)
	output, err := Filter(input, predicate)

	if err != nil {
		t.Error(err)
	}

	counter := 1
	for outputValue := range output {
		if outputValue != counter {
			t.Errorf("Filter failed, got %d, want %d", outputValue, counter)
		}
		counter++
	}
}