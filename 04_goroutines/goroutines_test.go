package main

import (
	"testing"
)

type TestCase struct {
	name         string
	inputValues  []int
	outputValues []int
	predicate    func(int) bool
}

func GetTestCases() []TestCase {
	return []TestCase{
		TestCase{
			"predicate returns true",
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			func(i int) bool { return true },
		},
		TestCase{
			"predicate returns false",
			[]int{1, 2, 3},
			[]int{},
			func(i int) bool { return false },
		},
	}
}
func TestFilter(t *testing.T) {
	testCases := GetTestCases()
	for _, testCase := range testCases {
		input := make(chan int, 5)

		for _, v := range testCase.inputValues {
			input <- v
		}
		close(input)

		output, err := Filter(input, testCase.predicate)
		if err != nil {
			t.Error(err)
		}

		for _, expectedValue := range testCase.outputValues {
			value := <-output
			if value != expectedValue {
				t.Errorf(
					"Filter failed when testing %s, got %d, want %d",
					testCase.name,
					value,
					expectedValue,
				)
			}
		}
	}

}
