package main

import (
	"fmt"
	"time"
)

// Define the DivideAndConquerable interface
type DivideAndConquerable interface {
	IsBasic() bool
	BaseFun() interface{}
	Decompose() []DivideAndConquerable
	Recombine(intermediateResults []interface{}) interface{}
	DivideAndConquer() interface{}
}

// Implement the DivideAndConquerable interface for a concrete type
type Fibonacci struct {
	input int
}

func (f Fibonacci) IsBasic() bool {
	return f.input == 0 || f.input == 1
}

func (f Fibonacci) BaseFun() interface{} {
	return f.input
}

func (f Fibonacci) Decompose() []DivideAndConquerable {
	return []DivideAndConquerable{
		Fibonacci{input: (f.input - 1)},
		Fibonacci{input: (f.input - 2)},
	}
}

func (f Fibonacci) Recombine(intermediateResults []interface{}) interface{} {
	result := 0
	for _, r := range intermediateResults {
		result += r.(int)
	}
	return result
}

func (f Fibonacci) DivideAndConquer() interface{} {
	if f.IsBasic() {
		return f.BaseFun()
	}
	subcomponents := f.Decompose()
	intermediateResults := make([]interface{}, len(subcomponents))
	ch := make(chan interface{}, len(subcomponents))

	// Launch goroutines for recursive calls
	for _, subcomponent := range subcomponents {
		go func(subcomponent DivideAndConquerable) {
			ch <- subcomponent.DivideAndConquer()
		}(subcomponent)
	}

	// Collect results from goroutines
	for i := 0; i < len(subcomponents); i++ {
		intermediateResults[i] = <-ch
	}

	close(ch)

	return f.Recombine(intermediateResults)
}

func main() {
	start := time.Now()
	result := Fibonacci{input: 40}.DivideAndConquer()
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Duration: %v\n", time.Since(start))
}
