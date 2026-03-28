package main

import (
	"sync"
)

func MonteCarlo(trials int, ng int) float64 {
	results := make(chan float64, trials)

	var wg sync.WaitGroup
	wg.Add(ng)
	chunk := trials / ng
	rem := trials % ng

	for i := 0; i < ng; i++ {
		n := chunk
		if i < rem { // adding the remainder to the first few goroutines
			n++
		}
		go func(numTrials int) {
			defer wg.Done()

			for j := 0; j < numTrials; j++ {
				// simulation trial
				// results <- value
			}
		}(n)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// aggregate results from channel
	// compute summary statistics
}