package fizzbuzz

import (
	"fmt"
	"testing"
)

var defaultResult = []string{"1", "2", "Fizz", "4", "Buzz",
	"Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}

var testcases = []struct {
	total  int64
	fizzAt int64
	buzzAt int64
	result []string
}{
	{total: 16, fizzAt: 3, buzzAt: 5, result: defaultResult},
}

func TestFizzBuzz(t *testing.T) {

	//var executeTests =
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
