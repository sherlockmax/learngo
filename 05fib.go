package main

import "fmt"

func main() {
	times := 100
	jobs := make(chan int, times)
	results := make(chan []int, times)

	go worker(jobs, results)

	for i := 1; i <= times; i++ {
		jobs <- i
	}
	close(jobs)

	for r := 0; r < times; r++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- []int) {
	for n := range jobs {
		results <- []int{n, fib(n)}
	}
}

// 斐波那契數列 (Fibonacci Number):每一項數字都是前面兩項數字的和
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
