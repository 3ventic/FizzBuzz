package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/3ventic/FizzBuzz/pkg/fizzbuzz"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGPIPE, syscall.SIGTERM)
	go func() {
		<-signals
		cancel()
	}()

	c, _ := fizzbuzz.FizzBuzz(ctx, map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		11: "Elevenish!",
	})
	for s := range c {
		fmt.Println(s)
	}
}
