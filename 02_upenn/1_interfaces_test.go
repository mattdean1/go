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
