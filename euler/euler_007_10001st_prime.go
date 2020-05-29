package euler

import "fmt"

// Problem 7: 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

func GetRankedPrime(rank int){
    count := 0
    rankedPrime := 2
    for i:=1;;i++{
        if IsPrime(i){
            count ++
            if count == rank{
                rankedPrime = i
                break
            }
        }
    }
    fmt.Printf("第%d个质数为%d.",rank, rankedPrime)
    fmt.Println()
}
