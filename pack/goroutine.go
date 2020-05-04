package pack

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

// Example3 calculating maximum average using sync.Mutex
func Example3() {

	var kk sync.Mutex
	workerCount := 1000
	do := make(chan []int, workerCount)

	var result int
	for i := 0; i < workerCount; i++ {

		go func() {
			arr := <-do
			average := avg(arr)
			kk.Lock()
			if result < average {
				result = average
			}
			kk.Unlock()
		}()

	}

	for i := 0; i < workerCount; i++ {
		arr := make([]int, 10)
		for j := 0; j < 10; j++ {
			arr[j] = rand.Int()
		}
		do <- arr
	}
	time.Sleep(2 * time.Second)
	fmt.Println(result)
}

// Example4 calculating maximum average using "MAP Reduced System"
func Example4() {

	workerCount := 1000
	do := make(chan []int, workerCount)

	result := make([]int, workerCount)
	for i := 0; i < workerCount; i++ {
		index := i
		go func() {
			arr := <-do
			average := avg(arr)
			if result[index] < average {
				result[index] = average
			}
		}()

	}

	for i := 0; i < workerCount; i++ {
		arr := make([]int, 10)
		for j := 0; j < 10; j++ {
			arr[j] = rand.Int()
		}
		do <- arr
	}

	var maxr int

	for _, each := range result {
		if maxr < each {
			maxr = each
		}
	}

	fmt.Println(maxr)
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
