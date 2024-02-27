package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
	for i, subcomponent := range subcomponents {
		intermediateResults[i] = subcomponent.DivideAndConquer()
	}
	return f.Recombine(intermediateResults)
}

func main() {
	start := time.Now()
	ch := make(chan interface{}, 2)

	wg.Add(1)
	go func(ch chan interface{}) {
		defer wg.Done()
		ch <- Fibonacci{input: 30}.DivideAndConquer()
	}(ch)

	wg.Add(1)
	go func(ch chan interface{}) {
		defer wg.Done()
		ch <- Fibonacci{input: 30}.DivideAndConquer()
	}(ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("Result: %v\n", result)
	}

	fmt.Printf("Duration: %v\n", time.Since(start))
}
