package main

/*Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]*/

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := range len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
