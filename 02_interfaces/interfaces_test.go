package main

import "testing"

// Implement sort.Interface
type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func TestIsPalindrome(t *testing.T) {
	shortPalindrome := StringSlice([]string{"a", "b", "a"})
	if !IsPalindrome(shortPalindrome) {
		t.Errorf("Expected %v to be a palindrome", shortPalindrome)
	}

	veryShortPalindrome := StringSlice([]string{"a"})
	if !IsPalindrome(veryShortPalindrome) {
		t.Errorf("Expected %v to be a palindrome", veryShortPalindrome)
	}

	notAPalindrome := StringSlice([]string{"a", "b", "c"})
	if IsPalindrome(notAPalindrome) {
		t.Errorf("Expected %v not to be a palindrome", notAPalindrome)
	}
}

func Sum(a, b int) int {
	return a + b
}
func TestFold(t *testing.T) {
	result := Fold([]int{1, 2, 3}, 0, Sum)
	if result != 6 {
		t.Errorf("Expected folding sum of [1,2,3] to be 6, actual %d", result)
	}

	emptySlice := []int{}
	initialValue := 0
	result = Fold(emptySlice, initialValue, Sum)
	if result != initialValue {
		t.Errorf("Expected folding sum of empty slice to be %d, actually %d", initialValue, result)
	}
}
