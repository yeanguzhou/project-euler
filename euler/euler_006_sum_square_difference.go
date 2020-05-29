package euler

import "fmt"

// Problem 6: Sum square difference
// The sum of the squares of the first ten natural numbers is, 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is, (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of
// the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers 
// and the square of the sum.

func SumSquareDiff(low int, top int) {
	sumOfSquares := 0
	squareOfSum := 0
	for i:=low;i<=top;i++{
		sumOfSquares += CalcSquare(i)
	}
	squareOfSum = CalcSquare(CalcSum(low, top))

	fmt.Printf("%d到%d的平方的和以及和的平方之差绝对值为%d.",low, top, IntAbs(sumOfSquares-squareOfSum))
	fmt.Println()
}

func IntAbs(num int) int {
	if num <=0 {
		return (-1) * num
	}else {
		return num
	}
}

func CalcSum(low int, top int) int {
	return (low + top) * (top - low + 1) /2
}

func CalcSquare(num int) int {
	return num * num
}
