package main

import (
	"testing"
)

func TestGetNums(t *testing.T) {
	testCases := []struct {
		testSet     string
		expectedAns int
	}{
		{"1eightcrcjcbdthreebscfpvznqfrj6", 16},
		{"ldfn2", 22},
		{"45", 45},
		{"9fivemksdnmgbvx", 99},
	}

	for _, tc := range testCases {
		result := getNums(tc.testSet)
		// TODO: compare the two arrays to check each element for test
		for index, item := range result {
			if item != tc.expectedAns {
				t.Errorf("Test Fail: expected %v, got %v for input %v, \n", tc.expectedAns, result, tc.testSet)
			}
		}
	}
}
