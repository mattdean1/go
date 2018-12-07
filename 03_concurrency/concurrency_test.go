package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestIOSum(t *testing.T) {
	var output bytes.Buffer
	input := strings.NewReader("1\n2\n")

	err := IOSum(input, &output)
	if err != nil {
		t.Error(err)
	}

	expectedSum := "3"
	actualSum := output.String()
	if actualSum != expectedSum {
		t.Errorf("IOSum failed, got: %s, want: %s", actualSum, expectedSum)
	}
}
