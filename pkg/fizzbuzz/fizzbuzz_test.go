package fizzbuzz_test

import (
	"context"
	"testing"

	"github.com/3ventic/FizzBuzz/pkg/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("errors with nil replacements", func(t *testing.T) {
		ctx := context.Background()
		c, err := fizzbuzz.FizzBuzz(ctx, nil)
		if err == nil {
			t.Errorf("expected non-nil error, got: %v", err)
		}
		if c != nil {
			t.Errorf("expected nil channel, got: %v", c)
		}
	})

	t.Run("succeeds with normal FizzBuzz", func(t *testing.T) {
		ctx := context.Background()
		replacements := map[int]string{
			3: "Fizz",
			5: "Buzz",
		}
		c, err := fizzbuzz.FizzBuzz(ctx, replacements)
		if err != nil {
			t.Errorf("expected nil error, got: %v", err)
		}
		if c == nil {
			t.Errorf("expected non-nil channel, got: %v", c)
		}
		expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19"}
		for _, e := range expected {
			value := <-c
			if value != e {
				t.Errorf("expected '%s' but got '%s'", e, value)
			}
		}
	})

	t.Run("succeeds with no replacements", func(t *testing.T) {
		ctx := context.Background()
		replacements := map[int]string{}
		c, err := fizzbuzz.FizzBuzz(ctx, replacements)
		if err != nil {
			t.Errorf("expected nil error, got: %v", err)
		}
		if c == nil {
			t.Errorf("expected non-nil channel, got: %v", c)
		}
		expected := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
		for _, e := range expected {
			value := <-c
			if value != e {
				t.Errorf("expected '%s' but got '%s'", e, value)
			}
		}
	})

	t.Run("succeeds with all primes under 19", func(t *testing.T) {
		ctx := context.Background()
		replacements := map[int]string{
			2:  "Two",
			3:  "Three",
			5:  "Five",
			7:  "Seven",
			11: "Eleven",
			13: "Thirteen",
			17: "Seventeen",
			19: "Nineteen",
		}
		c, err := fizzbuzz.FizzBuzz(ctx, replacements)
		if err != nil {
			t.Errorf("expected nil error, got: %v", err)
		}
		if c == nil {
			t.Errorf("expected non-nil channel, got: %v", c)
		}
		expected := []string{"1", "Two", "Three", "Two", "Five", "TwoThree", "Seven", "Two", "Three", "TwoFive", "Eleven", "TwoThree", "Thirteen", "TwoSeven", "ThreeFive", "Two", "Seventeen", "TwoThree", "Nineteen"}
		for _, e := range expected {
			value := <-c
			if value != e {
				t.Errorf("expected '%s' but got '%s'", e, value)
			}
		}
	})
}
