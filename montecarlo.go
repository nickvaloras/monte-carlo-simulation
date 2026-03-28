package main

import (
	"sync"
	"math/rand"
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
				// portfolio starts at $100
				value := 100.00
				for day := 0; day < 252; day++ { 
					// simulate 252 trading days with daily return of N(0.001, 0.02)
						r := 0.001 + 0.02 * rand.NormFloat64()
						value *= (1 + r)
				}
				// send simulated results through the channel
				results <- value
			}
		}(n)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// aggregate results from channel
	// compute summary statistics
	
	return 0.0 // temp
}