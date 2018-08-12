package euler

import (
    "fmt"
    "math"
    "reflect"
    "errors"
)

// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

// 判断obj是否在target中，target支持的类型array, slice, map
func Contains(obj interface{}, target interface{}) (bool, error) {
    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true, nil
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true, nil
        }
    }

    return false, errors.New("not in array")
}

// 质因数分解
func PrimeFactorization(number int)[]int{
    var factorList []int
    originNum := number

    var factor int
    for factor = 2;factor<=number;{
        if number % factor == 0 && IsPrime(factor){
            number = number/factor
            factorList = append(factorList, factor)
        }else {
            factor++
        }
    }
    fmt.Print("质因数分解：",originNum, "=")
    actualProduct := 1
    for index,value := range factorList{
       actualProduct *= value
       if index == len(factorList)-1{
           fmt.Print(value)
       }else {
           fmt.Print(value, "*")
       }
    }
    if actualProduct == originNum{
       fmt.Print(", 经过验算，结果正确!")
    }else {
       fmt.Print(", 经过验算，结果错误！验算值为",actualProduct)

    }
    fmt.Println()
    return factorList
}

// 获取质因数
func GetPrimeFactors(number int)  (primeFactorList []int,maxPrimeFactor int){
    var factorList []int
    var factor int
    originNum := number

    for factor = 2;factor<=number;{
        if number % factor == 0 && IsPrime(factor){
            number = number/factor
            if isIn,_:=Contains(factor, factorList);!isIn{
                factorList = append(factorList, factor)
            }
        }else {
            factor++
        }
    }
    maxPrimeFactor = factorList[len(factorList)-1]
    fmt.Print(originNum,"的质因数包括:", factorList, ", 最大质因数为",maxPrimeFactor, "。" )
    fmt.Println()
    return factorList, maxPrimeFactor
}

// 质数判断
func IsPrime(number int) bool {
    st := math.Sqrt(float64(number))
    isPm := true
    if number == 1{
        return false
    }
    if number == 2{
        return true
    }
    for i := 2;i <= int(st);i++{
        if number % i == 0{
            isPm = false
        }
    }
    return isPm
}