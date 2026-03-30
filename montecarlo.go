package main

import (
	"sync"
	"math/rand"
	"math"
)

// Create structure to return results as table (w/ mean, stddev, 95% CI)
type Result struct {
	Mean   float64
	StdDev float64
	CILow  float64
	CIHigh float64
}

func MonteCarlo(trials int, ng int) Result {
	if trials <= 0 || ng <= 0 {
		return Result{}
	}
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
	values := make([]float64, 0, trials)
	for v := range results {
		values = append(values, v)
	}
	n := float64(len(values))
	if n == 0 {
		return Result{}
	}

	// compute mean
	var sum float64
	for _, v := range values {
		sum += v
	}
	mean := sum / n

	// compute standard deviation
	var stdDev float64
	if n > 1 {
		var sqDiff float64
		for _, v := range values {
			diff := v - mean
			sqDiff += diff * diff
		}
		stdDev = math.Sqrt(sqDiff / (n - 1))
	}

	// construct 95% confidence interval
	margin := 1.96 * (stdDev / math.Sqrt(n))
	ciLow := mean - margin
	ciHigh := mean + margin
	
	// Use created table structure to return results

	return Result{
		Mean:   mean,
		StdDev: stdDev,
		CILow:  ciLow,
		CIHigh: ciHigh,
	}
}