package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	n := os.Args[1]
	i, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	fmt.Println(fib(i))
}
