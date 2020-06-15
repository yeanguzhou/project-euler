package main

import (
    "fmt"
    "time"

    "github.com/yeanguzhou/project-euler/euler"
)

func main()  {
    t0 := time.Now()

    euler.LatticePaths()

    t1 := time.Since(t0)
    fmt.Printf("\n程序总耗时%s.", t1)
}