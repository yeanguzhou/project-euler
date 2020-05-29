package euler

import "fmt"

// Problem 14: Longest Collatz sequence
// The following iterative sequence is defined for the set of positive integers:
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
// NOTE: Once the chain starts the terms are allowed to go above one million.


func LongestCollatzSequence() {
    maxChainLen := 1
    //fmt.Println(CollatzChain(13))
    chainMap := make(map[int]int)
    for i:=1;i< 1000000;i++{
       chainLen := CollatzChain(i)
       //fmt.Println(chainLen)
       chainMap[i] = chainLen
       if chainLen > maxChainLen{
          maxChainLen = chainLen
       }
    }

    for k, v := range chainMap{
        if v == maxChainLen{
            fmt.Println(k, v)
        }
    }
}

func CollatzChain(num int) int{
    chainLen := 0
    for next:=num;;next=CollatzRule(next){
        chainLen ++
        if next == 1{
            break
        }
    }
    return chainLen
}

func CollatzRule(num int) int {
    next := 1
	if num%2 == 0 {
	    next = num/2
	} else {
	    next = 3*num + 1
	}
	return next
}
