package main

import (
	"fmt"
	"time"
)

func main() {

	// ----------------------
	// Time Complexity - O(1)
	timeStart := time.Now()
	findSum(9999999999)
	timeEnd := time.Now()

	fmt.Println("Find Sum - Time taken: ", timeEnd.Sub(timeStart).Nanoseconds(), "ns")

	// ----------------------
	// Time Complexity - O(n)
	timeStart = time.Now()
	findSumIterate(9999999999)
	timeEnd = time.Now()

	fmt.Println("Find Sum Iterate - Time taken: ", timeEnd.Sub(timeStart).Milliseconds(), "ms")
}

// ----------------------
// PRIVATE FUNCTIONS
// ----------------------S

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumIterate(n int) int {

	sum := 0

	for i := 0; i < n; i++ {
		sum += i
	}

	return sum
}
