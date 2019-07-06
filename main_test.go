package main

import "testing"

var testcases = []struct {
	input    string
	expected string
	nums     []int
}{
	{"", "No such command found!", []int{}},
	{"Hello", "Hello World!", []int{}},
	{"SumXY", "3", []int{1, 2}},                                           // sum with 2 value
	{"SumXY", "6", []int{1, 2, 3}},                                        // sum with 3 value
	{"SumXY", "-2", []int{1, 2, -5}},                                      // sum with negative value
	{"MultiplyXY", "2", []int{1, 2}},                                      // multiplication with 3 value
	{"MultiplyXY", "-24", []int{2, 3, 4, -1}},                             // multiplication with negative value
	{"MultiplyXY", "0", []int{2, 3, 4, 0}},                                // multiplication with zero value
	{"FindFirstNPrime", "2", []int{1}},                                    // Find first prime number
	{"FindFirstNPrime", "2, 3, 5, 7", []int{4}},                           // Find first 4 prime number
	{"FindFirstNPrime", "2, 3, 5, 7, 11, 13, 17, 19, 23", []int{9}},       // Find first 9 prime number
	{"FindFirstNFibonacci", "0, 1, 1, 2", []int{4}},                       // Find first 4 fibonacci sequence
	{"FindFirstNFibonacci", "0, 1, 1, 2, 3, 5, 8, 13, 21, 34", []int{10}}, // Find first 10 fibonacci sequence
}

func TestExecFuncByName(t *testing.T) {
	for _, tt := range testcases {
		t.Run(tt.input, func(t *testing.T) {
			s := ExecFuncByName(tt.input, tt.nums...)
			if s != tt.expected {
				t.Errorf("got %v, expected %v", s, tt.expected)
			}
		})
	}
}
