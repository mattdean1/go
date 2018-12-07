package main

import (
	"testing"
	"strings"
	"bytes"
)

func TestIOSum(t *testing.T) {
	var output bytes.Buffer
	input := strings.NewReader("1\n2\n")

	IOSum(input, &output)

	expectedSum := "3"
	actualSum := output.String()
	if actualSum != expectedSum {
		t.Errorf("IOSum failed, got: %s, want: %s", actualSum, expectedSum)
	}
}