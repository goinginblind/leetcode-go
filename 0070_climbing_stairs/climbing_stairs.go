package main

import "fmt"

//	func climbStairs(n int) int {
//		if n <= 2 {
//			return n
//		}
//		return climbStairs(n-1) + climbStairs(n-2)
//	}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(7))
}
