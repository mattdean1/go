package main

import (
	"fmt"
	"testing"
)

type FilterTestCase struct {
	name         string
	inputValues  []int
	outputValues []int
	predicate    func(int) bool
}

func GetFilterTestCases() []FilterTestCase {
	return []FilterTestCase{
		FilterTestCase{
			"predicate returns true",
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			func(i int) bool { return true },
		},
		FilterTestCase{
			"predicate returns false",
			[]int{1, 2, 3},
			[]int{},
			func(i int) bool { return false },
		},
	}
}
func TestFilter(t *testing.T) {
	testCases := GetFilterTestCases()
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


func TestConcurrentRetry(t *testing.T) {
	tasks := make([](func () (string, error)), 10)
	for i := 0; i < 10; i++ {
		i := i // assign inside loop so we can use the value
		tasks[i] = func() (string, error) {
			return fmt.Sprintf("task %d", i), nil
		}
	}

	results := ConcurrentRetry(tasks, 1, 1)
	for result := range results {
		fmt.Printf("result %d: %s \n", result.id, result.result)
	}
}
