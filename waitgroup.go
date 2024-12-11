package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this worker as done when it finishes
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each worker
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers are done.")
}
