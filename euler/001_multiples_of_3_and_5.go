package euler

import "fmt"

// Problem 1: Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func GetMultiples()  {
    var multipleList []int
    var sum int
    for i :=1; i <1000; i++ {
        if i % 3 == 0 || i % 5 == 0{
            sum += i
            multipleList = append(multipleList, i)
        }
    }
    fmt.Println(multipleList)
    fmt.Println(sum)
}
