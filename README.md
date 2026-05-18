# Parallel Monte Carlo Simulation for Portfolio Valuation

A concurrent Monte Carlo simulation in Go that estimates the expected one-year value of a financial portfolio. Trials are distributed across goroutine workers and aggregated via channels, with benchmarks measuring parallel speedup and statistical convergence across trial counts up to 100 million.

## Overview

Monte Carlo methods estimate quantities by averaging over many independent random trials — a class of problem that parallelizes naturally. This project implements the pattern in Go using goroutines, channels, and `sync.WaitGroup` to coordinate workers and aggregate results.

**Estimated quantities per run:**
- Mean final portfolio value
- Standard deviation
- 95% confidence interval

## Model

| Parameter | Value |
|-----------|-------|
| Initial investment | $100 |
| Time horizon | 252 trading days (1 year) |
| Daily return distribution | Normal(μ = 0.001, σ = 0.02) |

Each trial simulates 252 daily returns drawn from the distribution above and compounds them to produce a final portfolio value.

## Architecture

- **Workers** (goroutines): each worker processes an independent chunk of trials
- **Channels**: workers send results to a central aggregator
- **WaitGroup**: ensures all workers finish before final statistics are computed

## Usage
go run montecarlo.go [trials] [workers]
Example:
go run montecarlo.go 1000000 8

## Benchmarks

Runtime measured across trial counts (10K → 100M) and worker counts (4, 6, 8).

### Real time by configuration

| Trials       | 4 Workers | 6 Workers | 8 Workers |
|--------------|-----------|-----------|-----------|
| 10,000       | 0.445s    | 0.350s    | 0.586s    |
| 100,000      | 0.504s    | 0.558s    | 0.540s    |
| 1,000,000    | 1.151s    | 1.372s    | 4.223s    |
| 10,000,000   | 7.940s    | 8.362s    | 8.528s    |
| 100,000,000  | 95.40s    | 80.38s    | 66.41s    |

### Speedup relative to 4 workers

| Trials       | 4 Workers | 6 Workers | 8 Workers |
|--------------|-----------|-----------|-----------|
| 10,000       | 1.00x     | 1.27x     | 0.76x     |
| 100,000      | 1.00x     | 0.90x     | 0.93x     |
| 1,000,000    | 1.00x     | 0.84x     | 0.27x     |
| 10,000,000   | 1.00x     | 0.95x     | 0.93x     |
| 100,000,000  | 1.00x     | 1.19x     | 1.44x     |

**Observations:**

- At small trial counts (≤1M), goroutine and channel overhead dominate the actual compute. Adding workers can *hurt* performance — visible in the 8-worker / 1M case running 4× slower than 4 workers.
- Parallel speedup only emerges clearly at 100M trials, where 8 workers achieve a 1.44× speedup over 4 workers. This is the regime where useful work per goroutine exceeds coordination cost.
- The non-monotonic behavior across mid-range trial counts is consistent with Amdahl's Law: the sequential aggregation step (computing summary statistics over collected results) caps achievable speedup, and that ceiling becomes binding earlier when per-worker workloads are small.

### Statistical convergence

| Trials       | Mean    | Std Dev | 95% CI Width |
|--------------|---------|---------|--------------|
| 10,000       | 128.91  | 41.76   | 1.63         |
| 100,000      | 128.48  | 41.72   | 0.52         |
| 1,000,000    | 128.60  | 41.78   | 0.17         |
| 10,000,000   | 128.61  | 41.82   | 0.05         |
| 100,000,000  | 128.64  | 41.85   | 0.02         |

The mean stabilizes around **$128.6** and the standard deviation around **$41.8** after only ~100K trials. The 95% CI width, however, shrinks by ~80× from 10K to 100M trials — consistent with the √N convergence rate expected of Monte Carlo estimators. In practice this means additional trials don't change the estimate, they just sharpen its precision.

## Tech Stack

- **Language:** Go
- **Concurrency primitives:** goroutines, channels, `sync.WaitGroup`
- **Domain:** quantitative finance, stochastic simulation, parallel computing

## Background

Built as a course project for CMDA 3634 (Concurrent & Parallel Computing) at Virginia Tech.

