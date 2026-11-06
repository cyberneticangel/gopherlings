// sync.WaitGroup counts jobs: Add(1) before starting each goroutine,
// Done() when it finishes, Wait() until the count reaches zero. With
// each goroutine sending its partial result on a shared channel, this
// is the fan-out/fan-in pattern.
package main

import (
	"fmt"
	"sync"
)

// parallelSum sums nums by splitting it into parts chunks and adding
// up each chunk in its own goroutine.
func parallelSum(nums []int, parts int) int {
	if parts <= 0 {
		return 0
	}
	size := (len(nums) + parts - 1) / parts

	partials := make(chan int, parts)
	var wg sync.WaitGroup
	for start := 0; start < len(nums); start += size {
		end := start + size
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			total := 0
			for _, v := range chunk {
				total += v
			}
			partials <- total
		}(nums[start:end])
	}
	wg.Wait()
	close(partials)

	sum := 0
	for partial := range partials {
		sum += partial
	}
	return sum
}

func main() {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}
	fmt.Println("sum with 4 goroutines:", parallelSum(nums, 4))
}
