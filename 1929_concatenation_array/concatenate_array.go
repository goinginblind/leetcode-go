package main

// [1,3,2,1] --> [1,3,2,1,1,3,2,1]

func getConcatenation(nums []int) []int {
	ans := make([]int, 0, len(nums)*2)

	ans = append(ans, nums...)
	ans = append(ans, nums...)

	return ans
}

func main() {
	arr := []int{1, 3, 2, 1}
	getConcatenation(arr)
}
