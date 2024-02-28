package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var inputNumbers = []int{2, 3, 4, 5}

// DivideAndConquerable defines the interface of the base code
type DivideAndConquerable interface {
	IsBasic() bool
	BaseFun() interface{}
	Decompose() []DivideAndConquerable
	Recombine(intermediateResults []interface{}) interface{}
	DivideAndConquer() interface{}
}

// Fibonacci defines the struct for computing Fibonacci sequence
type Fibonacci struct {
	input int
}

// IsBasic returns the basic cases of the Fibonacci sequence
func (f Fibonacci) IsBasic() bool {
	return f.input == 0 || f.input == 1
}

// BaseFun checks if it is a base case of the Fibonacci sequence
func (f Fibonacci) BaseFun() interface{} {
	return f.input
}

// Decompose returns the subcomponents of the Fibonacci sequence
func (f Fibonacci) Decompose() []DivideAndConquerable {
	return []DivideAndConquerable{
		Fibonacci{input: (f.input - 1)},
		Fibonacci{input: (f.input - 2)},
	}
}

// Recomibe returns the sum of the intermidate results of the Fibonacci sequence
func (f Fibonacci) Recombine(intermediateResults []interface{}) interface{} {
	result := 0
	for _, r := range intermediateResults {
		result += r.(int)
	}
	return result
}

// DivideAndConquer calculates the Fibonacci sequence
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
	var end time.Time
	start := time.Now()

	ch := make(chan interface{}, len(inputNumbers))

	for _, inputNumber := range inputNumbers {
		wg.Add(1)
		go func(ch chan interface{}, inputNumber int) {
			defer wg.Done()
			ch <- Fibonacci{input: inputNumber}.DivideAndConquer()
		}(ch, inputNumber)
	}

	go func() {
		wg.Wait()
		end = time.Now()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("Result: %v\n", result)
	}

	fmt.Printf("Duration: %v\n", end.Sub(start))
}
