package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Return the result
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to signal no more jobs

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}
