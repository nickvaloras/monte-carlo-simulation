# CMDA3634 Project 1 - Monte Carlo Simulation

## Overview

This project implements a parallel Monte Carlo simulation in Go to estimate the expected value of a financial portfolio over one year.

The simulation models daily returns as a normal distribution and uses goroutines and channels to parallelize independent trials.

## How to Run Code

go run montecarlo.go [trials] [workers]

## Model Description

- Initial investment: $100
- Period: 252 trading days (1 year)
- Daily return:
  
  R ~ N(μ = 0.001, σ = 0.02)

Each trial simulates one year of returns and outputs the final portfolio value.

## Method

Our Monte Carlo simulation is used to estimate:

- Average (mean) final portfolio value
- Standard deviation
- 95% confidence interval

The simulation is parallelized using:

- Goroutines (workers)
- Channels (aggregation)
- WaitGroups (synchronization)

Each worker runs a chunk of trials independently and sends results to a central aggregator through channels.

## AI Use
We did not use AI on this assignment. We relied on examples in class and similar exercises to design a simple parallel implementation of the Monte Carlo simulation.