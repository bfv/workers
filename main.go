package main

import "fmt"

const workers = 4

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// first we fill up our jobs channel. If size permits all jobs can be put channel at one go
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	// start our workers
	for w := 0; w < workers; w++ {
		go worker(jobs, results)
	}

	// read the results
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	// loop & wait for values on the jobs channel
	for n := range jobs {
		results <- fib(n)
	}
}

// inefficient algoritm to generate some load
func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
