package euler

import (
	"errors"
	"fmt"
)

// Problem 15: Lattice paths
// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20×20 grid?

// Solution 1:
// For a m x m grid, the route length n=2m. Then the answer is C(m, n).
//  C(n,m) = n!/((n-m)!*m!) = (2m)!/(m!*m!) = (m+m)(m+(m-1))(m+(m-2))...(m+2)(m+1)/((m(m-1)(m-2))))
// = (m+m)/m*(m+(m-1))/(m-1)*...*(m+2)/2*(m+1)/1
// = product{m+i/i}
// Time complexity：O(5m+2)=O(m)
func CombinationForLatticePaths(m int64) int64 {
	result := int64(1)
	for i := int64(1); i <= m; i++ {
		result = result * (m + i) / i
	}
	return result
}

// Solution 2:
// C(n, m) = C(n-1, m) + C(n-1, m-1)
// Solve recursively, but it's a low effective algorithm if n is large.
// Time complexity: O(2C(n, m)-1) =O(C(n,m))
func CombinationRecursively(n int64, m int64) (int64, error) {
	if m > n || m < 0 {
		return -1, errors.New("index out of range")
	}
	if n-m < m {
		return CombinationRecursively(n, n-m)
	}
	if m == 0 {
		return 1, nil
	} else if m == n {
		return 1, nil
	} else {
		c1, err1 := CombinationRecursively(n-1, m)
		c2, err2 := CombinationRecursively(n-1, m-1)
		if err1 != nil {
			return -1, err1
		}
		if err2 != nil {
			return -1, err2
		}
		return c1 + c2, nil
	}
}

// The Lattice Paths Problem is a combinatorial problem.
// Through a 20×20 grid, there are 20 horizontal moves and 20 vertical moves.
// Mark a horizontal move as 1 and a vertical move as 0. Then this problem transfers to the problem：
// how many 40 digit numbers are there which includes 20 digits of 1 and 20 digits of 0.
// And of course the answer is C(20, 40).

// 组合问题。在每一个路径中，都会有20次横向和20次纵向移动。
// 如果把横向移动记为1，纵向移动记为0，那么问题就变成一个40位数字中有20个1和20个0的组合数。所以答案就是C（20,40）。
func LatticePaths() {
	fmt.Println(CombinationForLatticePaths(20))
	//fmt.Println(CombinationRecursively(6, 2))
}
