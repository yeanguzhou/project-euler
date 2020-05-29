package euler

import "fmt"

// Problem 9: Special Pythagorean triplet
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func SpecialPythagoreanTriplet()  {
	for a:=1;a<=332;a++{
		for b:=a+1;b<(1000-a-b);b++{
			if CalcSquare(a)+CalcSquare(b)==CalcSquare(1000-a-b){
				fmt.Println(a,b, 1000-a-b, a*b*(1000-a-b))
			}
		}
	}
}
