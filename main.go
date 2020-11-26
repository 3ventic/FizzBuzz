package main

import (
	"context"
	"fmt"
	"time"

	"github.com/3ventic/FizzBuzz/pkg/fizzbuzz"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, _ := fizzbuzz.FizzBuzz(ctx, map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		11: "Elevenish!",
	})
	for s := range c {
		fmt.Println(s)
	}
}
