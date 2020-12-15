package fizzbuzz

import (
	"fmt"
	"testing"
)

var default16Result = []string{"1", "2", "Fizz", "4", "Buzz",
	"Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}

var reverse16Result = []string{"1", "2", "Buzz", "4", "Fizz",
	"Buzz", "7", "8", "Buzz", "Fizz", "11", "Buzz", "13", "14", "FizzBuzz", "16"}

var multiples16Result = []string{"1", "2", "3", "Fizz", "5",
	"6", "7", "FizzBuzz", "9", "10", "11", "Fizz", "13", "14", "15", "FizzBuzz"}

var testcases = []struct {
	total  int64
	fizzAt int64
	buzzAt int64
	result []string
}{
	{total: 16, fizzAt: 3, buzzAt: 5, result: default16Result},
	{total: 16, fizzAt: 5, buzzAt: 3, result: reverse16Result},
	{total: 16, fizzAt: 4, buzzAt: 8, result: multiples16Result},
	{total: 5, fizzAt: 2, buzzAt: 2, result: []string{"1", "FizzBuzz", "3", "FizzBuzz", "5"}},
	{total: 3, fizzAt: 5, buzzAt: 9, result: []string{"1", "2", "3"}},
	{total: 0, fizzAt: 5, buzzAt: 6, result: []string{}},
	{total: 6, fizzAt: -2, buzzAt: -3, result: []string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz"}},
}

func TestFizzBuzz(t *testing.T) {

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("total: %d, fizzAt: %d, buzzAt: %d", tc.total, tc.fizzAt, tc.buzzAt),
			func(t *testing.T) {
				res := FizzBuzz(tc.total, tc.fizzAt, tc.buzzAt)
				if len(res) != len(tc.result) {
					t.Errorf("Expected result length %d, got length %d", len(tc.result), len(res))
				}
				for i, expected := range tc.result {
					if res[i] != expected {
						t.Errorf("Expected %s got %s at result position %d", expected, res[i], i)
					}
				}
			})
	}
}
