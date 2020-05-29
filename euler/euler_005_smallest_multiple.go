package euler

import (
	"fmt"
)

// Problem 5: Smallest multiple
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

var factorMap map[int]int  // factorMap[factor] = count
var tempFactorMap map[int]int

func GetSmallestMultiple(low int, top int) int {
	// 规律：能被指定范围内所有整数都整除的最小数为其中所有整数的质因数分解后，所有质因数按最高幂次之积。
	// 10 = 2*5 = 2^1 * 5^1
	// 9 = 3*3 = 3^2
	// 8 = 2*2*2 = 2^3
	// 7 = 7 = 7^1
	// 6 = 2 * 3 = 2^1 * 3^1
	// 5 = 5
	// 4 = 2 * 2 = 2^2
	// 3 = 3
	// 2 = 2
	// 1 = 1

	factorMap = make(map[int]int)

	for i:=low;i<=top;i++{
		tempFactorMap = make(map[int]int)
		for _,v := range PrimeFactorization(i){ // 质因数分解后，逐个计数
			tempFactorMap[v]++
			if tempFactorMap[v] >= factorMap[v]{
				factorMap[v] = tempFactorMap[v]
			}
		}
	}
	smallestMultiPle := 1
	for factor := range factorMap{
		smallestMultiPle *= IntPow(factor, factorMap[factor])
	}
	fmt.Printf("能被%d~%d之间的所有整数都整除的最小正整数为%d.",low, top,smallestMultiPle)
	fmt.Println()
	return int(smallestMultiPle)
}

func IntPow(x int, y int) int {
	pow := 1
	for i:=0;i<y;i++{
		pow *= x
	}
	return pow
}
