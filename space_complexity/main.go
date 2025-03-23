package main

import (
	"fmt"
	"runtime"
)

func main() {

	var mBefore, mAfter runtime.MemStats

	/*
		The code below is intended to empirically demonstrate the concept.
		If you encounter performance or memory constraints, use a smaller input size
		to clearly observe the relationship between input size and memory allocation.
	*/

	// ----------------------
	// Space Complexity - smaller input size
	runtime.ReadMemStats(&mBefore)
	findSumIterateStore(999)
	runtime.ReadMemStats(&mAfter)

	fmt.Println("Find Sum Iterate (Smaller Input Size) - Mem Alloc: ", mAfter.Alloc-mBefore.Alloc, "bytes")

	// ----------------------
	// Space Complexity - larger input size
	runtime.ReadMemStats(&mBefore)
	findSumIterateStore(9999)
	runtime.ReadMemStats(&mAfter)

	fmt.Println("Find Sum Iterate (Larger Input Size) - Mem Alloc: ", mAfter.Alloc-mBefore.Alloc, "bytes")

}

// ----------------------
// PRIVATE FUNCTIONS
// ----------------------

func findSumIterateStore(n int) int {

	result := make([]int, 0, n)

	sum := 0
	for i := 0; i < n; i++ {
		sum += i
		result = append(result, i)
	}

	return sum
}
