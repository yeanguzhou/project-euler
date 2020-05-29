package euler

import "fmt"

// Problem 10: Summation of primes
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

func GetSumOfPrimes(number int)  {
    sum := 0
    for i:=2;i<number;i++{
        if IsPrime(i){
            sum += i
        }
    }
    fmt.Printf("%d以内的所有质数之和为%d.", number, sum)
    fmt.Println()
}
