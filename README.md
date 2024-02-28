# Calculate Fibonacci Sequence in Golang using Recursion

Used at the university to challenge different programming languages. 
Therefore we had to implement the same basic code structure and play around with multi threading and caching.

## Project structure

- concurrency-base &rarr; first implementation with basic struct from the professor
- concurrency-caching &rarr; caches the calculations
- goroutines-caching &rarr; combination of goroutines-multiple-inputs and the caching feature from concurrency-caching
- goroutines-multiple-inputs &rarr; calculates from multiple number inputs in parallel
- goroutines-on-recursion &rarr; sends every recursion to a new goroutine (does not calculate in parallel)
- goroutines-on-subcomponents &rarr; calculates subcomponents in parallel