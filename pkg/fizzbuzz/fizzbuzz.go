package fizzbuzz

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// FizzBuzz returns a channel that emits numbers starting from 1 until the context is cancelled with multipleReplacements done
// multipleReplacements should be a map, where each multiple to be replaced is the key and its replacement is the corresponding value,
// e.g. map[int]string{3: "Fizz", 5: "Buzz"} for the classic example of this problem.
func FizzBuzz(ctx context.Context, multipleReplacements map[int]string) (<-chan string, error) {
	if multipleReplacements == nil {
		return nil, errors.New("multipleReplacements cannot be nil")
	}
	multiples := make([]int, 0, len(multipleReplacements))
	maxLen := 0
	for m, v := range multipleReplacements {
		multiples = append(multiples, m)
		maxLen += len(v)
	}
	c := make(chan string)
	go fizzBuzz(ctx, multiples, multipleReplacements, maxLen, c)
	return c, nil
}

func fizzBuzz(ctx context.Context, multiples []int, multipleReplacements map[int]string, maxLen int, c chan<- string) {
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			close(c)
			return
		default:
			var sb strings.Builder
			sb.Grow(maxLen)
			for _, m := range multiples {
				if i%m == 0 {
					sb.WriteString(multipleReplacements[m])
				}
			}
			if sb.Len() == 0 {
				c <- strconv.Itoa(i)
			} else {
				c <- sb.String()
			}
		}
	}
}
