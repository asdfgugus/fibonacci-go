package main

import (
	"fmt"
	"sync"
	"time"
)

var inputNumber int = 10
var wg sync.WaitGroup
var mu sync.Mutex

// Define the DivideAndConquerable interface
type DivideAndConquerable interface {
	IsBasic() bool
	BaseFun() interface{}
	Decompose() []DivideAndConquerable
	Recombine(intermediateResults []interface{}) interface{}
	DivideAndConquer(ch chan interface{})
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

func (f Fibonacci) DivideAndConquer(ch chan interface{}) {
	defer wg.Done()
	if f.IsBasic() {
		ch <- f.BaseFun()
		return
	}
	subcomponents := f.Decompose()
	intermediateResults := make([]interface{}, len(subcomponents))
	for i, subcomponent := range subcomponents {
		wg.Add(1)
		go func(i int, sub DivideAndConquerable) {
			sub.DivideAndConquer(ch)
		}(i, subcomponent)
	}

	for i := range subcomponents {
		intermediateResults[i] = <-ch
	}

	mu.Lock()
	ch <- f.Recombine(intermediateResults)
	mu.Unlock()
}

func main() {
	start := time.Now()
	ch := make(chan interface{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		Fibonacci{input: 30}.DivideAndConquer(ch)
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("Result: %v\n", result)
	}

	fmt.Printf("Duration: %v\n", time.Since(start))
}
