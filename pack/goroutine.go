package pack

import (
	"fmt"
	"sync"
)

// Example1 is basic implementation of goroutines with sync.WaitGroup
func Example1() {
	var wg sync.WaitGroup
	s := []int{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			PrintSlice(s)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Example1 is completed\n")
}

// Example2 is basic implementation of passing function in channel inside go routine
func Example2() {
	doFunc := make(chan func([]int) int)
	s := []int{1, 2, 3, 4}
	workerCount := 10
	result := make([]int, workerCount)

	for i := 0; i < workerCount; i++ {
		index := i
		go func() {
			f := <-doFunc
			average := f(s)
			if result[index] < average {
				result[index] = average
			}

		}()
	}
	doFunc <- avg
	doFunc <- avg
	doFunc <- avg
	doFunc <- avg
	doFunc <- avg
	doFunc <- sum
	doFunc <- sum
	doFunc <- sum
	doFunc <- sum
	doFunc <- sum

	fmt.Println(result)
	fmt.Println("Example2 is completed\n")

}

// sum to calculate sum of slice
func sum(s []int) int {
	sum := 0
	for _, each := range s {
		sum += each
	}
	return sum
}

// avg of a slice
func avg(s []int) int {
	return sum(s) / len(s)
}

// PrintSlice to print a slice
func PrintSlice(a []int) {
	fmt.Println(a)
}
