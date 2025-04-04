package main

import "fmt"

func main() {

	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	oddNumbers := removeEvenNumbersList(input)

	fmt.Println("RESULT ODD NUMBERS: ", oddNumbers)

	reverse := reverseList(input)

	fmt.Println("RESULT REVERSE LIST: ", reverse)

	mvZerosA := moveZerosA(input)
	fmt.Println("RESULT MOVE ZEROS TO THE END LIST - BETTER: ", mvZerosA)

	mvZerosB := moveZerosB(input)
	fmt.Println("RESULT MOVE ZEROS TO THE END - WORSE: ", mvZerosB)

	missingNUmber := missingNumber(input)
	fmt.Println("RESULT MISSING NUMBER: ", missingNUmber)

}

func removeEvenNumbersList(input []int) (output []int) {

	for i, value := range input {
		if input[i]%2 != 0 {
			output = append(output, value)
		}
	}

	return output
}

func reverseList(input []int) []int {

	output := make([]int, len(input))
	copy(output, input)

	start := 0
	end := len(input) - 1

	for start < end {
		output[start], output[end] = output[end], output[start]
		start++
		end--
	}

	return output
}

func moveZerosA(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)

	j := 0
	arrLen := len(input)

	for i := 0; i < arrLen; i++ {
		if output[i] != 0 && output[j] == 0 {
			output[i], output[j] = output[j], output[i]
		}
		if output[j] != 0 {
			j++
		}
	}

	return output
}

func moveZerosB(input []int) []int {

	output := make([]int, len(input))
	copy(output, input)

	arrLen := len(output)

	for i := 0; i < arrLen; i++ {
		if output[i] == 0 {
			output = append(output[:i], output[i+1:]...)
			output = append(output, 0)
			i -= 1
			arrLen -= 1
		}
	}

	return output

}

func missingNumber(input []int) int {

	arrLen := len(input)
	sum := arrLen * (arrLen + 1) / 2

	for _, value := range input {
		sum -= value
	}

	return sum
}
