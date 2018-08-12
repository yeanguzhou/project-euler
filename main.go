package main

import (
    "project-euler/euler"
    "time"
    "fmt"
)

func main()  {
    t0 := time.Now()

    euler.LongestCollatzSequence()

    t1 := time.Since(t0)
    fmt.Println()
    fmt.Printf("程序总耗时%s.", t1)
}