package main

import (
	"sync"
)

func MonteCarlo(trials int, ng int) float64 {
	results := make(chan float64, trials)

	var wg sync.WaitGroup
	wg.Add(ng)
	chunk := trials / ng

	for i := 0; i < ng; i++ {
		go func(n int) {
			defer wg.Done()

			for j := 0; j < n; j++ {
				// simulation trial
				// results <- value
			}
		}(chunk)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// aggregate results from channel
	// compute summary statistics
}