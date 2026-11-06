// sync.WaitGroup tracks a group of goroutines and waits for them all:
//
//	var wg sync.WaitGroup
//	wg.Add(1)           // one more job to wait for
//	go func() {
//		defer wg.Done() // my job is finished
//		// ... work ...
//	}()
//	wg.Wait()           // block until every job is done
//
// Together with a shared channel this is the fan-out/fan-in pattern:
// split the work across goroutines (fan-out), let each send its
// partial result, and combine the results in the caller (fan-in).
//
// Task: implement parallelSum. Split nums into `parts` contiguous
// chunks, run one goroutine per chunk, and have each goroutine send
// its partial sum on a shared channel. Wait for them all with a
// sync.WaitGroup, then add up the partials and return the total.
package main

import "fmt"

// parallelSum sums nums by splitting it into parts chunks and adding
// up each chunk in its own goroutine.
func parallelSum(nums []int, parts int) int {
	return 0 // TODO: fan out one goroutine per chunk, then sum the partials they send back
}

func main() {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}
	fmt.Println("sum with 4 goroutines:", parallelSum(nums, 4))
}

// I AM NOT DONE
