package euler

import (
    "fmt"
    "strconv"
)

// Problem 4: Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func IsPalindrome(number int)  bool{
    numStr := strconv.Itoa(number)
    numLen := len(numStr)
    isPa := true
    if numLen % 2 == 0{
        for i:=0;i<=(numLen/2)-1;i++{
            if numStr[i] != numStr[numLen-1-i]{
                isPa = false
                break
            }
        }
    } else {
        for i:=0;i<=((numLen+1)/2)-1;i++{
            if numStr[i] != numStr[numLen-1-i]{
                isPa = false
                break
            }
        }
    }
    return isPa
}

func GetLargestPalindrome(digitLen int)  {
    largestPa := 0
    product := 1
    leftFactor := 1
    rightFactor := 9

    minStr := "1"
    maxStr := "9"
    if digitLen < 2{
        minStr = "0"
        maxStr = "9"
    }else {
        for i:=1;i<digitLen;i++{
            minStr = minStr + "0"
            maxStr = maxStr + "9"
        }
    }
    min, _ := strconv.Atoi(minStr)
    max, _ := strconv.Atoi(maxStr)

    for i:=min;i<=max;i++{
        for j:=min;j<=max;j++{
            product = i*j
            if IsPalindrome(product){
                //fmt.Println(product)
                if product >= largestPa{
                    largestPa = product
                    leftFactor = i
                    rightFactor = j
                }
            }
        }
    }
    fmt.Printf("由两个%d位数相乘所得回文的最大值为%d=%d*%d.",digitLen, largestPa, leftFactor, rightFactor)
}