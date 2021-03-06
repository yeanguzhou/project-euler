package euler

import "fmt"

// Problem 12: Highly divisible triangular number
// The sequence of triangle numbers is generated by adding the natural numbers.
// So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
// Let us list the factors of the first seven triangle numbers:
// 1: 1
// 3: 1,3
// 6: 1,2,3,6
//10: 1,2,5,10
//15: 1,3,5,15
//21: 1,3,7,21
//28: 1,2,4,7,14,28
// We can see that 28 is the first triangle number to have over five divisors.
// What is the value of the first triangle number to have over five hundred divisors?

var factMap map[int]int  // factorMap[factor] = count
var tempMap map[int]int

func HighDivTriNum() {
    for rank:=1;;rank++{
        tn := TriangleNumber(rank)
        count := CountDivisors(tn)
        if count > 500{
            fmt.Printf("第%d个三角形数为%d，它有%d个因数.", rank, tn, count)
            break
        }
    }
}

// 计算第rank个三角形数
func TriangleNumber(rank int) int {
	var triNum int
	for i := 0; i <= rank; i++ {
		triNum += i
	}
	return triNum
}

// 计算因数总个数
func CountDivisors(num int) int {
    // 质因数分解后，因数个数取决于各质因数幂次。
    // 计算方式96=2^5*3^1, (5+1) * (1+1)，
    // 即n=x^a*y^b*z^c...,则因数个数为(a+1)(b+1)(c+1)...
    factMap = make(map[int]int)
    tempMap = make(map[int]int)

    for _,v := range PrimeFactorization(num){
        tempMap[v]++
        if tempMap[v] >= factMap[v]{
            factMap[v] = tempMap[v]
        }
    }

    divCounter := 1
    for _, c := range factMap{
        divCounter = divCounter * (c+1)
    }
    return divCounter

    // 低效方式：逐个判断
    //
    //count := 0
    //for i:=1;i<=num;i++ {
    //    if num % i == 0{
    //        count ++
    //    }
    //}
    //return count
}
